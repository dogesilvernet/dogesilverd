#!/bin/bash
rm -rf /tmp/dogesilverd-temp

NUM_CLIENTS=128
dogesilverd --devnet --appdir=/tmp/dogesilverd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
DOGESILVERD_PID=$!
DOGESILVERD_KILLED=0
function killDogesilverdIfNotKilled() {
  if [ $DOGESILVERD_KILLED -eq 0 ]; then
    kill $DOGESILVERD_PID
  fi
}
trap "killDogesilverdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $DOGESILVERD_PID

wait $DOGESILVERD_PID
DOGESILVERD_EXIT_CODE=$?
DOGESILVERD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilverd exit code: $DOGESILVERD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVERD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
