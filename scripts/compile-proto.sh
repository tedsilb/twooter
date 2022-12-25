#!/bin/bash

cd proto

protoc \
    --go_out=twooterpb \
    --go_opt=paths=source_relative \
    --go-grpc_out=twooterpb \
    --go-grpc_opt=paths=source_relative \
    *.proto
