package client

import (
	"context"
	"time"

	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the dogesilverwalletd server, and returns the client instance
func Connect(address string) (pb.DogesilverwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("dogesilverwallet daemon is not running on Starbase, start it with `dogesilverwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewDogesilverwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
