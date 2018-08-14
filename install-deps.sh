#! /bin/sh

# Python
pip install grpcio-tools
pip install googleapis-common-protos

# Go
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
