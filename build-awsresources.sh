#!/bin/bash

echo "package metrics"

echo "var awsResources = []string{"

while IFS='' read -r line || [[ -n "$line" ]]; do
    echo "\"$line\"," 
done < "$1"

echo "}"
