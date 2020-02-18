#!/bin/sh

ARTIFACT=$1
MAIN_PATH=$2

set -eu

if [ -z "$ARTIFACT" ]
then
    echo 'missing artifact name'
    exit 1
fi

if [ -z "$MAIN_PATH" ]
then
    MAIN_PATH="./cmd"
fi

export GO111MODULE=on

go get -v ./...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $ARTIFACT $MAIN_PATH