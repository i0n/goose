#!/bin/bash
#
# This script builds the application from source.
set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd $DIR

# Build!
echo "--> Building OSX Binary..."
go build \
    -o bin/osx/goose
cp bin/osx/goose $(echo $GOPATH | sed -e 's/:.*//g')/bin

echo "--> Building Linux Binary..."
GOOS=linux GOARCH=amd64 go build \
    -o bin/linux/goose
cp bin/linux/goose $(echo $GOPATH | sed -e 's/:.*//g')/bin

echo "--> Building Windows Binary..."
GOOS=windows GOARCH=386 go build \
    -o bin/windows/goose.exe
cp bin/windows/goose.exe $(echo $GOPATH | sed -e 's/:.*//g')/bin
