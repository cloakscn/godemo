#!/bin/bash

os() {
  if uname -s | grep Darwin; then
    return 2
  elif uname -s | grep Linux; then
    return 1
  else
    #    windows
    return 3
  fi
}

linux() {
  return
}

windows() {
  return
}

darwin() {
  return
}

GOOS=${1:-default}
GOARCH=${2:-default}
os
OS_FLAG=$?

# 执行系统独有的配置
if [ $OS_FLAG = 1 ]; then
  linux
elif [ $OS_FLAG = 2 ]; then
  darwin
elif [ $OS_FLAG = 3 ]; then
  windows
fi

if [ "$GOOS" != "default" ]; then
  if [ "$GOARCH" != "default" ]; then
    echo "set environment variable GOOS=${GOOS} GOARCH=${GOARCH}"
    export GOOS=${GOOS}
    export GOARCH=${GOARCH}
    export CGO_ENABLED=0
  else
    echo "error input, goos or goarch is error input"
    return
  fi
else
  export CGO_ENABLED=1
fi

echo "go build ${GOOS} ${GOARCH}"
go build main.go
