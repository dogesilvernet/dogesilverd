#!/bin/bash
rm -rf /tmp/dogesilverd-temp

dogesilverd --devnet --appdir=/tmp/dogesilverd-temp --profile=6061 --loglevel=debug &
DOGESILVERD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $DOGESILVERD_PID

wait $DOGESILVERD_PID
DOGESILVERD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilverd exit code: $DOGESILVERD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVERD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
