#!/bin/bash

# Generate aws.go from list of aws resources in awsresources.txt
./build-awsresources.sh awsresources.txt > aws.go

# Format aws.go file
go fmt aws.go
