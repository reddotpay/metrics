#!/bin/bash
#
# Generate README

set -ex

PKG_LIST=$(go list ./... | grep -v /vendor/)

printf "" > README.md

# Create a coverage file for each package
if [ "${package}" != "git.reddotpay.com/core-service/$1" ]; then
    godocdown "${package}" >> README.md;
fi

printf "\n\n" >> README.md
printf '```' >> README.md
printf "\n" >> README.md
gocloc . >> README.md
printf '```' >> README.md