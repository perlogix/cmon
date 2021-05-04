#!/bin/sh

mkdir -p /etc/perlogix/cmon

mv -f /etc/perlogix/cmon/cmon /usr/bin/cmon
chmod -f 0755 /usr/bin/cmon

if [ ! -e /etc/perlogix/cmon/cmon.yaml ]; then
  cp -f /etc/perlogix/cmon/cmon.yaml.orig /etc/perlogix/cmon/cmon.yaml
fi

if [ "$(pgrep systemd -c)" -ge 2 ]; then
  INIT="systemd"
else
  INIT="other"
fi

if [ "$INIT" = "systemd" ]; then
  cp -f /etc/perlogix/cmon/cmon.service /etc/systemd/system/cmon.service
  systemctl daemon-reload
  systemctl restart cmon
fi

if [ "$INIT" = "other" ]; then
  cp -f /etc/perlogix/cmon/cmon.init /etc/init.d/cmon
  chmod -f 0755 /etc/init.d/cmon
  chkconfig cmon on
  /etc/init.d/cmon restart
fi
