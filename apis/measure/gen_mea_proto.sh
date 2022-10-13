#!/usr/bin/env bash

# Copyright (c) 2019-2029 Dunyu All Rights Reserved.
#
# Author      : yangping
# Email       : ping.yang@wengold.net
# Version     : 1.0.0
# Description :
#   Generate measure rgpc server and client.
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2022/10/213  yangping       New version
# -------------------------------------------------------------------

bin=`dirname "$0"`
bin=`cd "$bin"; pwd`

cd ./${bin}
protoc --go_out=. --go-grpc_out=. ./proto/mea.proto

