#!/bin/bash
#
# Generate README

set -ex

godocdown "${package}" > README.md;

printf "\n\n" >> README.md
printf '```' >> README.md
printf "\n" >> README.md
gocloc . >> README.md
printf '```' >> README.md