#!/bin/sh

set -eu

export GO111MODULE=on

if [ -z "$LOG_LEVEL" ]
then
    export LOG_LEVEL=Debug
fi

if [ -z "$DB_HOST" ]
then
    export DB_HOST=localhost
fi

if [ -z "$DB_PORT" ]
then
    export DB_PORT=27017
fi


if [ -z "$DB_DATABASE" ]
then
    export DB_DATABASE=Instance
fi



go get -t -v ./...
go test ./... -timeout=2m -parallel=4
