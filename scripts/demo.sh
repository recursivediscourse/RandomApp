#! /bin/sh

set -o errexit
set -o nounset

cd /app
pwd
ls -la
echo Hello world, from a script file!

# Notes for future reference
# `set -o errexit` makes sure your script exits when your command fails
# `set -o nounset` exits the script when it tries to call undeclared vars
# `set -o pipefail` 
# 
