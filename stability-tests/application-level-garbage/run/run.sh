#!/bin/bash
rm -rf /tmp/dogesilverd-temp

dogesilverd --devnet --appdir=/tmp/dogesilverd-temp --profile=6061 --loglevel=debug &
DOGESILVER_PID=$!
DOGESILVER_KILLED=0
function killDogesilverdIfNotKilled() {
    if [ $DOGESILVER_KILLED -eq 0 ]; then
      kill $DOGESILVER_PID
    fi
}
trap "killDogesilverdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DOGESILVER_PID

wait $DOGESILVER_PID
DOGESILVER_KILLED=1
DOGESILVER_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Dogesilverd exit code: $DOGESILVER_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DOGESILVER_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
