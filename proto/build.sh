#!/bin/bash

if [[ ! -d ./googleapis ]]; then
  git clone https://github.com/googleapis/googleapis.git
fi

if [[ ! -d ./googleapis ]]; then
  echo "missing repo: https://github.com/googleapis/googleapis.git"
  exit 1
fi

echo "build api proto with grpc-gateway"
protoc -I./ -I./googleapis --go_out=.. --go-grpc_out=.. api/*.proto
cd ../protobuf/api
go mod init github.com/binbin6363/icuc/protobuf/api
go mod tidy
cd -

echo "build im/app proto"
protoc --go_out=.. --go-grpc_out=.. im/app/*.proto
cd ../protobuf/im/app
go mod init github.com/binbin6363/icuc/protobuf/im/app
go mod tidy
cd -
