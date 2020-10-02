#!/bin/sh

if [ "$(pgrep systemd -c)" -ge 2 ]; then
    INIT="systemd"
else
    INIT="other"
fi

if [ "$INIT" = "systemd" ]; then
    systemctl stop yeti-discover
    rm -rf /etc/systemd/system/yeti-discover.service /usr/lib/yeticloud/yeti-discover
    systemctl daemon-reload
fi
if [ "$INIT" = "other" ]; then
    /etc/init.d/yeti-discover stop
    chkconfig yeti-discover off
    rm -rf /etc/init.d/yeti-discover /usr/lib/yeticloud/yeti-discover
fi
