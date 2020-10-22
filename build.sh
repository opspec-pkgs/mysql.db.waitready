#!/bin/sh

go mod vendor
env GOOS=linux GOARCH=amd64 go build -o waitready
rm -rf vendor