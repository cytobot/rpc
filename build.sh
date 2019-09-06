#!/bin/bash

sudo apt install -y libprotoc-dev libprotobuf-dev protobuf-compiler

export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:/$GOROOT/bin:/$GOPATH/bin

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

find . -name '*.proto' -type f | while read line; do
    protoc -I $(dirname "$line") $line --go_out=plugins=grpc:$(basename "${line%.*}")
done