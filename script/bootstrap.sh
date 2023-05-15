#!/bin/bash

WEB_LOG_FILE=./log/web.log
USER_LOG_FILE=./log/user.log
PLAY_LOG_FILE=./log/play.log
TICKET_LOG_FILE=./log/ticket.log
STUDIO_LOG_FILE=./log/studio.log
ORDER_LOG_FILE=./log/order.log

# 创建日志文件并清空内容
echo "" >${WEB_LOG_FILE}
echo "" >${USER_LOG_FILE}
echo "" >${PLAY_LOG_FILE}
echo "" >${TICKET_LOG_FILE}
echo "" >${STUDIO_LOG_FILE}
echo "" >${ORDER_LOG_FILE}

# 启动程序，并将输出重定向到日志文件
go build -o ttms_web ../internal/web
./ttms_web >>${WEB_LOG_FILE} 2>&1 &
go build -o ttms_user ../internal/user
./ttms_user >>${USER_LOG_FILE} 2>&1 &
go build -o ttms_play ../internal/play
./ttms_play >>${PLAY_LOG_FILE} 2>&1 &
go build -o ttms_ticket ../internal/ticket
./ttms_ticket >>${TICKET_LOG_FILE} 2>&1 &
go build -o ttms_studio ../internal/studio
./ttms_studio >>${STUDIO_LOG_FILE} 2>&1 &
go build -o ttms_order ../internal/order
./ttms_order >>${ORDER_LOG_FILE} 2>&1 &
