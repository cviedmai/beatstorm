#!/bin/bash
if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

FROM=9095
TO=9096
if grep --quiet 9096 /opt/beatstorm/config.json; then
  FROM=9096
  TO=9095
fi

function rollback {
  sed -i "s/$1/$2/g" /opt/beatstorm/config.json
  sed -i "s/$1/$2/g" /opt/nginx/sites-available/beatstorm
}

sed -i "s/$FROM/$TO/g" /opt/beatstorm/config.json
if [ $? -ne 0 ]; then
  echo "failed to sed beatstorm config"
  exit 1
fi

sed -i "s/$FROM/$TO/g" /opt/nginx/sites-available/beatstorm
if [ $? -ne 0 ]; then
  rollback $TO $FROM
  echo "failed to sed nginx config"
  exit 1
fi

beatstorm_PID=`(cat /opt/beatstorm/beatstorm.pid)`
cd /home/USERNAME/go/src/beatstorm/
export GOROOT=/opt/go
export GOPATH=/home/USERNAME/go/
/opt/go/bin/go build -a app/beatstorm.go
mv /opt/beatstorm/beatstorm /opt/beatstorm/beatstorm.old
mv beatstorm /opt/beatstorm/
kill -INT $beatstorm_PID
sleep 10
/etc/init.d/nginx reload
