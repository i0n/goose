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

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Build!
echo "--> Building OSX Binary..."
go build \
    -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
    -v \
    -o bin/osx/goose
cp bin/osx/goose $(echo $GOPATH | sed -e 's/:.*//g')/bin

echo "--> Building Linux Binary..."
GOOS=linux GOARCH=amd64 go build \
    -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
    -v \
    -o bin/linux/goose
cp bin/linux/goose $(echo $GOPATH | sed -e 's/:.*//g')/bin
