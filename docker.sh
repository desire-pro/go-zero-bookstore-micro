#!/bin/sh

cd ./api
pwd
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build bookstore.go
bookstore="apirpc:0.0.1"
echo "開始生成Image: ${bookstore}"
docker build . -t ${bookstore}

cd ./../rpc/add
pwd
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build add.go
addrpc="addrpc:0.0.1"
echo "開始生成Image: ${addrpc}"
docker build . -t ${addrpc}

cd ./../check
pwd
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build check.go
checkrpc="checkrpc:0.0.1"
echo "開始生成Image: "${checkrpc}
docker build . -t ${checkrpc}

cd ./../change
pwd
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build change.go
changerpc="changerpc:0.0.1"
echo "開始生成Image: "${changerpc}
docker build . -t ${changerpc}