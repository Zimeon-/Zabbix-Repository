#! /bin/sh
SERVICE=$1
TIMEOUT=25
SERVICE_STATUS=`echo | sudo /etc/init.d/$1 status 2>/dev/null`
if [[ $SERVICE_STATUS == *"is running..."* ]]
then
        SERVICE_STATUS=1
else
        SERVICE_STATUS=0
fi
echo ${SERVICE_STATUS}
