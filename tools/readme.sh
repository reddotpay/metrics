#!/bin/bash
#
# Generate README

set -ex

PKG_LIST=$(go list ./... | grep -v /vendor/)

# Create a coverage file for each package
for package in ${PKG_LIST}; do
    if [ "${package}" != "git.reddotpay.com/core-service/$1" ]; then
        godocdown "${package}" >> README.md;
    fi
done ;

printf "\n\n" >> README.md
printf '```' >> README.md
printf "\n" >> README.md
gocloc . >> README.md
printf '```' >> README.md