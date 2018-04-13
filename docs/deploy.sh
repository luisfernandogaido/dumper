#!/usr/bin/env bash
file=/var/www/html/dumper/deploy/dumper
if [ -e "$file" ]; then
    systemctl stop dumper.service
    mv "$file" /var/www/html/dumper
    chmod 0774 /var/www/html/dumper/dumper
    systemctl start dumper.service
fi