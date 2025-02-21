#!/bin/bash
rm -rf /tmp/dogesilver-temp

dogesilver --simnet --appdir=/tmp/dogesilver-temp --profile=6061 &
DOGESILVERD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $DOGESILVERD_PID

wait $DOGESILVERD_PID
DOGESILVERD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilver exit code: $DOGESILVERD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVERD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
