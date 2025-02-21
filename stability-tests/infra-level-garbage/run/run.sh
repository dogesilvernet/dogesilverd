#!/bin/bash
rm -rf /tmp/dogesilverd-temp

dogesilverd --devnet --appdir=/tmp/dogesilverd-temp --profile=6061 &
DOGESILVER_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DOGESILVER_PID

wait $DOGESILVER_PID
DOGESILVER_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilverd exit code: $DOGESILVER_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVER_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
