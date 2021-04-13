#! /bin/sh

set -o errexit
set -o nounset
set -o pipefail

pwd
ls -la
echo Hello world, from a script file!
go run cmd/main.go

