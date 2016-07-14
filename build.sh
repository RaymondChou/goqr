#!/bin/sh

go build main.go

# Compile statically linking (avoiding need of a fat image)
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create image
docker build -t loggi/goqr .
