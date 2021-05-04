#!/bin/sh

if [ "$(pgrep systemd -c)" -ge 2 ]; then
  INIT="systemd"
else
  INIT="other"
fi

if [ "$INIT" = "systemd" ]; then
  systemctl stop cmon
  rm -rf /etc/systemd/system/cmon.service /etc/perlogix/cmon /usr/bin/cmon
  systemctl daemon-reload
fi

if [ "$INIT" = "other" ]; then
  /etc/init.d/cmon stop
  chkconfig cmon off
  rm -rf /etc/init.d/cmon /etc/perlogix/cmon /usr/bin/cmon
fi
