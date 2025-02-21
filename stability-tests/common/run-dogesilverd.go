package common

import (
	"fmt"
	"github.com/dogesilvernet/dogesilverd/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunDogesilverdForTesting runs dogesilverd for testing purposes
func RunDogesilverdForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	dogesilverdRunCommand, err := StartCmd("DOGESILVER",
		"dogesilverd",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Dogesilverd started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := dogesilverdRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Dogesilverd closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := dogesilverdRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Dogesilverd stopped")
	}
}
