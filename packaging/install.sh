#!/bin/sh

mkdir -p /etc/yeticloud /usr/lib/yeticloud/yeti-discover

mv -f /usr/lib/yeticloud/yeti-discover/yeti-discover /usr/bin/yeti-discvoer
chmod -f 0755 /usr/bin/yeti-discover

if [ ! -e /etc/yeticloud/yeti-discover.yaml ]; then
  mv -f /usr/lib/yeticloud/yeti-discover/yeti-discover.yaml /etc/yeticloud/yeti-discover.yaml
fi

if [ "$(pgrep systemd -c)" -ge 2 ]; then
  INIT="systemd"
else
  INIT="other"
fi

if [ "$INIT" = "systemd" ]; then
  cp -f /usr/lib/yeticloud/yeti-discover/yeti-discover.service /etc/systemd/system/yeti-discover.service
  systemctl daemon-reload
  systemctl restart yeti-discover
fi

if [ "$INIT" = "other" ]; then
  cp -f /usr/lib/yeticloud/yeti-discover/yeti-discover.init /etc/init.d/yeti-discover
  chmod -f 0755 /etc/init.d/yeti-discover
  chkconfig yeti-discover on
  /etc/init.d/yeti-discover restart
fi
