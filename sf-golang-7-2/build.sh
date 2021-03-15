#!/usr/bin/env sh

rm -rf ./build
go build -ldflags "-s -w" -o ./build/sf-golang-7-2-darwin .
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./build/sf-golang-7-2-linux-arm64 .
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/sf-golang-7-2-win64.exe .
