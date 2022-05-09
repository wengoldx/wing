#!/usr/bin/env bash

# Copyright (c) 2018-2028 Dunyu All Rights Reserved.
#
# Author      : https://www.wengold.net
# Email       : support@wengold.net
#
# Prismy.No | Date       | Modified by. | Description
# -------------------------------------------------------------------
# 00001       2020/05/08   yangping       New version
# 00002       2020/08/16   yangping       Support for windows
# -------------------------------------------------------------------

bin=`dirname "$0"`
bin=`cd "$bin"; pwd`

source ${bin}/exports.sh
${bin}/daemon.sh start ${SERVICE_APP_NAME}

# waiting for 5 seconds later
sleep 5

# start browser and fullscreen to show dashbord page
# chromee --start-fullscreen "http://${SERVICE_HOST_PORT}/${SERVICE_APP_NAME}/"
chromee --kiosk "http://${SERVICE_HOST_PORT}/${SERVICE_APP_NAME}/"
