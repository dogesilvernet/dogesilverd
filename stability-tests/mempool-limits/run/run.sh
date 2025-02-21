#!/bin/bash

APPDIR=/tmp/dogesilverd-temp
DOGESILVERD_RPC_PORT=29587

rm -rf "${APPDIR}"

dogesilverd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${DOGESILVERD_RPC_PORT}" --profile=6061 &
DOGESILVERD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${DOGESILVERD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $DOGESILVERD_PID

wait $DOGESILVERD_PID
DOGESILVERD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilverd exit code: $DOGESILVERD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVERD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
