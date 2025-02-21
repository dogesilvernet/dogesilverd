package server

import (
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dogesilvernet/dogesilverd/version"

	"github.com/dogesilvernet/dogesilverd/domain/consensus/model/externalapi"

	"github.com/dogesilvernet/dogesilverd/util/txmass"

	"github.com/dogesilvernet/dogesilverd/util/profiling"

	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/pb"
	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/keys"
	"github.com/dogesilvernet/dogesilverd/domain/dagconfig"
	"github.com/dogesilvernet/dogesilverd/infrastructure/network/rpcclient"
	"github.com/dogesilvernet/dogesilverd/infrastructure/os/signal"
	"github.com/dogesilvernet/dogesilverd/util/panics"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDogesilverwalletdServer

	rpcClient           *rpcclient.RPCClient // RPC client for ongoing user requests
	backgroundRPCClient *rpcclient.RPCClient // RPC client dedicated for address and UTXO background fetching
	params              *dagconfig.Params
	coinbaseMaturity    uint64 // Is different from default if we use testnet-11

	lock                            sync.RWMutex
	utxosSortedByAmount             []*walletUTXO
	mempoolExcludedUTXOs            map[externalapi.DomainOutpoint]*walletUTXO
	nextSyncStartIndex              uint32
	keysFile                        *keys.File
	shutdown                        chan struct{}
	forceSyncChan                   chan struct{}
	startTimeOfLastCompletedRefresh time.Time
	addressSet                      walletAddressSet
	txMassCalculator                *txmass.Calculator
	usedOutpoints                   map[externalapi.DomainOutpoint]time.Time
	firstSyncDone                   atomic.Bool

	isLogFinalProgressLineShown bool
	maxUsedAddressesForLog      uint32
	maxProcessedAddressesForLog uint32
}

// MaxDaemonSendMsgSize is the max send message size used for the daemon server.
// Currently, set to 100MB
const MaxDaemonSendMsgSize = 100_000_000

// Start starts the dogesilverwalletd server
func Start(params *dagconfig.Params, listen, rpcServer string, keysFilePath string, profile string, timeout uint32) error {
	initLog(defaultLogFile, defaultErrLogFile)

	defer panics.HandlePanic(log, "MAIN", nil)
	interrupt := signal.InterruptListener()

	if profile != "" {
		profiling.Start(profile, log)
	}

	log.Infof("Version %s", version.Version())
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return (errors.Wrapf(err, "Error listening to TCP on %s", listen))
	}
	log.Infof("Listening to TCP on %s", listen)

	log.Infof("Connecting to a node at %s...", rpcServer)
	rpcClient, err := connectToRPC(params, rpcServer, timeout)
	if err != nil {
		return (errors.Wrapf(err, "Error connecting to RPC server %s", rpcServer))
	}
	backgroundRPCClient, err := connectToRPC(params, rpcServer, timeout)
	if err != nil {
		return (errors.Wrapf(err, "Error making a second connection to RPC server %s", rpcServer))
	}

	log.Infof("Connected, reading keys file %s...", keysFilePath)
	keysFile, err := keys.ReadKeysFile(params, keysFilePath)
	if err != nil {
		return (errors.Wrapf(err, "Error reading keys file %s", keysFilePath))
	}

	err = keysFile.TryLock()
	if err != nil {
		return err
	}

	dagInfo, err := rpcClient.GetBlockDAGInfo()
	if err != nil {
		return nil
	}

	coinbaseMaturity := params.BlockCoinbaseMaturity
	if dagInfo.NetworkName == "dogesilver-testnet-11" {
		coinbaseMaturity = 1000
	}

	serverInstance := &server{
		rpcClient:                   rpcClient,
		backgroundRPCClient:         backgroundRPCClient,
		params:                      params,
		coinbaseMaturity:            coinbaseMaturity,
		utxosSortedByAmount:         []*walletUTXO{},
		mempoolExcludedUTXOs:        map[externalapi.DomainOutpoint]*walletUTXO{},
		nextSyncStartIndex:          0,
		keysFile:                    keysFile,
		shutdown:                    make(chan struct{}),
		forceSyncChan:               make(chan struct{}),
		addressSet:                  make(walletAddressSet),
		txMassCalculator:            txmass.NewCalculator(params.MassPerTxByte, params.MassPerScriptPubKeyByte, params.MassPerSigOp),
		usedOutpoints:               map[externalapi.DomainOutpoint]time.Time{},
		isLogFinalProgressLineShown: false,
		maxUsedAddressesForLog:      0,
		maxProcessedAddressesForLog: 0,
	}

	log.Infof("Read, Starbase start syncing the wallet...")
	spawn("serverInstance.syncLoop", func() {
		err := serverInstance.syncLoop()
		if err != nil {
			printErrorAndExit(errors.Wrap(err, "Starbase response error syncing the wallet"))
		}
	})

	grpcServer := grpc.NewServer(grpc.MaxSendMsgSize(MaxDaemonSendMsgSize))
	pb.RegisterDogesilverwalletdServer(grpcServer, serverInstance)

	spawn("grpcServer.Serve", func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			printErrorAndExit(errors.Wrap(err, "Error serving gRPC"))
		}
	})

	select {
	case <-serverInstance.shutdown:
	case <-interrupt:
		const stopTimeout = 2 * time.Second

		stopChan := make(chan interface{})
		spawn("gRPCServer.Stop", func() {
			grpcServer.GracefulStop()
			close(stopChan)
		})

		select {
		case <-stopChan:
		case <-time.After(stopTimeout):
			log.Warnf("Starbase could not gracefully stop: timed out after %s", stopTimeout)
			grpcServer.Stop()
		}
	}

	return nil
}

func printErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	os.Exit(1)
}
