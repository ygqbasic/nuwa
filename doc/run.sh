#!/bin/bash

case $1 in 
	start)
		nohup ./nuwa 2>&1 >> info.log 2>&1 /dev/null &
		echo "service start..."
		sleep 1
	;;
	stop)
		killall nuwa
		echo "service stop..."
		sleep 1
	;;
	restart)
		killall nuwa
		sleep 1
		nohup ./nuwa 2>&1 >> info.log 2>&1 /dev/null &
		echo "service restart..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
