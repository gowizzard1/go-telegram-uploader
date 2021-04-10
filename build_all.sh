#!/usr/bin/env bash

pushd .
HERE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
# echo $HERE
# cd $HERE
go get ./...
./go_multi_arch_build.sh github.com/techieIsaac/go-telegram-uploader ./build 'windows/amd64|linux/amd64|linux/arm/5'
popd
