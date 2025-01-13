#!/bin/bash

# exit if error
set -e

# create bin if does not exists
mkdir -p ./bin

# compile the app
go build -o ./bin/pmp ./cmd/b2b-service-pmp/main.go

echo "Build completed successfully."