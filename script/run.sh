#!/usr/bin/env bash

# template
cd service/template
rm -f template
go build -o template ./cmd/

echo "Run on :8080"
./template -port=8080


