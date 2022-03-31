#!/bin/bash

# Script for building the go zip file for AWS

GOOS=linux GOARCH=amd64 go build -o main main.go zip main.zip main