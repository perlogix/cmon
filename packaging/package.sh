#!/bin/bash

mkdir -p /usr/lib/yeticloud/yeti-discover
cp -f ./* /usr/lib/yeticloud/yeti-discover/
rm -f /usr/lib/yeticloud/yeti-discover/package.sh
fpm -t deb -s dir -n yeti-discover -v "$1" -a amd64 -p "yeti-discover-$1-amd64.deb" --license GPLv3 --vendor YetiCloud -m hello@yeticloud.com --url "https://github.com/yeticloud/yeti-discover" --description "YetiCloud yeti-discover binary distribution" --after-install ./install.sh
fpm -t rpm -s dir -n yeti-discover -v "$1" -a amd64 -p "yeti-discover-$1-amd64.rpm" --license GPLv3 --vendor YetiCloud -m hello@yeticloud.com --url "https://github.com/yeticloud/yeti-discover" --description "YetiCloud yeti-discover binary distribution" --after-install ./install.sh
