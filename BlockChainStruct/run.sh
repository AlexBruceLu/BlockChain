#!/bin/bash
rm block
rm -rf blockChain.db
rm -rf blockChain.db.lock
go build -o block *.go
./block

