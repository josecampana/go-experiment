#!/bin/bash

# exit if error
set -e

# check if go.mod exists
if [ ! -f go.mod ]; then
  echo "go.mod not found. Please ensure you are in the correct directory."
  exit 1
fi

# install dependencias
echo "tidying up the module dependencies..."
go mod tidy

echo "downloading module dependencies..."
go mod download

echo "dependencies installed successfully."