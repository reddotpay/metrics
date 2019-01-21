#!/bin/bash

echo "package metrics"

echo "var awsResources = map[string]int{"

while IFS='' read -r line || [[ -n "$line" ]]; do
    echo "\"$line\": 1," 
done < "$1"

echo "}"

echo "// AWSResource defines aws resources"
echo "type AWSResource string"

while IFS='' read -r line || [[ -n "$line" ]]; do
    s=$(echo $line | tr -d ' ' | tr -d '.' | tr -d '-')
    echo "// $s defines $line resource"
    echo "const $s AWSResource = \"$line\"" 
done < "$1"
