#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -o ogt-ags-64.exe main.go
GOOS=windows GOARCH=386 go build -o ogt-ags-32.exe main.go
