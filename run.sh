#!/usr/bin/env bash

case $1 in
	start)
		sudo nohup server >/dev/null 2>&1 &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall server
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall server
		sleep 1
		sudo nohup server  >/dev/null 2>&1 &
		echo "服务已重启..."
		sleep 1
	;;
	*)
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac