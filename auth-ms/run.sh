#!/usr/bin/env bash
go test ./...
go build ./cmd/main.go
./main -env=.env