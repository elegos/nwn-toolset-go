#!/usr/bin/env bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin

# Aurora library
go test -v -cover aurora/...
