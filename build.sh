#!/bin/bash

mkdir -p output/bin output/conf output/log
cp script/bootstrap.sh output
chmod +x output/bootstrap.sh

find conf/ -type f ! -name  "*_local.*" | xargs -I{} cp {} output/conf/

RUN_NAME="main_xx"
go build -o output/bin/${RUN_NAME} -v
