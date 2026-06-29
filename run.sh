#!/bin/bash

if ! command -v go >/dev/null 2>&1; then
    echo "Go is not installed"
    echo " First instlal go"
    exit 1
fi

echo "Building the CATOoOoO ..."
go build -buildvcs=false -o cato

if [ $? -eq 0 ]; then
    echo "Running Cato..."
    ./cato
else
    echo "hmmm, smth bad happened smth is wrong"
    exit 1
fi