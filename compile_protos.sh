#!/bin/bash

yes | rm -R ./protocol/generated/*
protoc --go_out=./protocol/ --go_opt=paths=import --go_opt=module=github.com/jamesrobb/ably-takehome/protocol \
--go-grpc_out=./protocol/ --go-grpc_opt=paths=import --go-grpc_opt=module=github.com/jamesrobb/ably-takehome/protocol \
protocol/protocol.proto