#!/bin/bash
export BS_ROOT=$(pwd)
export BS_ENV=test

for package in `go list ./...`; do
  go test -i $package
  go test $package
  if [[ $? != 0 ]] ; then exit $?; fi
done

#coffee tests/integration/run.coffee
