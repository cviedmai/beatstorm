#!/bin/bash
beatstorm_ROOT=/opt/beatstorm beatstorm_ENV=production /opt/beatstorm/beatstorm . > /var/log/opt/beatstorm.log 2>&1 &
