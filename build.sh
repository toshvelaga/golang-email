#!/bin/bash

# Script for building the go zip file for AWS

GOARCH=amd64 GOOS=linux go build main.go
zip main.zip main 