#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/pb2emmy ./cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/pb2emmy ./cmd/main.go