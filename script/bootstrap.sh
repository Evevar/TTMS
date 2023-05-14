#!/bin/bash

WEB_LOG_FILE=/var/log/TTMS/web.log
USER_LOG_FILE=/var/log/TTMS/user.log
PLAY_LOG_FILE=/var/log/TTMS/play.log
TICKET_LOG_FILE=/var/log/TTMS/ticket.log
STUDIO_LOG_FILE=/var/log/TTMS/studio.log
ORDER_LOG_FILE=/var/log/TTMS/order.log

# 创建日志文件并清空内容
echo "" >${WEB_LOG_FILE}
echo "" >${USER_LOG_FILE}
echo "" >${PLAY_LOG_FILE}
echo "" >${TICKET_LOG_FILE}
echo "" >${STUDIO_LOG_FILE}
echo "" >${ORDER_LOG_FILE}

# 启动程序，并将输出重定向到日志文件
go build ../internal/web &
./web >>${WEB_LOG_FILE} 2>&1 &
go build ../internal/user &
./user >>${USER_LOG_FILE} 2>&1 &
go build ../internal/play &
./play >>${PLAY_LOG_FILE} 2>&1 &
go build ../internal/ticket &
./ticket >>${TICKET_LOG_FILE} 2>&1 &
go build ../internal/studio &
./studio >>${STUDIO_LOG_FILE} 2>&1 &
go build ../internal/order &
./order >>${ORDER_LOG_FILE} 2>&1 &
