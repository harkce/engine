#!/bin/bash

echo "mode: atomic" > coverage.out

PACKAGES=`go list ./...`
SKIPPED=(
  github.com/harkce/engine/app/engine~
)

EXIT_CODE=0

for PKG in $PACKAGES; do
  if [[ "${SKIPPED[@]}" =~ "${PKG}~" ]]; then
    go test -v $PKG; __EXIT_CODE__=$?
  else
    go test -v -coverprofile=profile.out -covermode=atomic $PKG; __EXIT_CODE__=$?
  fi

  if [ "$__EXIT_CODE__" -ne "0" ]; then
    EXIT_CODE=$__EXIT_CODE__
  fi

  if [ -f profile.out ]; then
    tail -n +2 profile.out >> coverage.out; rm profile.out
  fi
done

exit $EXIT_CODE

