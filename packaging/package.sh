#!/bin/bash

if [[ "$1" == "" ]]; then
  echo "First argument empty, need version"
  exit 1
fi

mkdir -p /usr/lib/yeticloud/yeti-discover

cp -f ./yeti-discover.service /usr/lib/yeticloud/yeti-discover/
cp -f ../yeti-discover /usr/lib/yeticloud/yeti-discover/

cat <<'EOF' >/usr/lib/yeticloud/yeti-discover/yeti-discover.yaml
# ElasticSearch DB
#host: localhost
#port: 9200

# ElasticSearch Index Name
# This can be anything, it could be aws, softlayer, prod, staging
#environment: dev

# Interval of agent runs in seconds
# Default is every twenty minutes
#interval: 1200

# Username if http-basic plugin is enabled
#username:

# Password if http-basic plugin is enabled
#password:

# https true enables HTTPS instead of HTTP)
#https: false

# Verify SSL for HTTP endpoints
#verify_ssl: true

# Public facing asset
#public: false

# Asset type
#asset_type:

# OpenSCAP XCCDF XML file path
#oscap_xccdf_xml: /usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml

# OpenSCAP Profile
#oscap_profile: xccdf_org.ssgproject.content_profile_cis
EOF

fpm -t deb -s dir -n yeti-discover -v "$1" -a amd64 -p "yeti-discover-$1-amd64.deb" --license GPLv3 --vendor YetiCloud -m hello@yeticloud.com --url "https://github.com/yeticloud/yeti-discover" --description "YetiCloud yeti-discover binary distribution" --after-install ./install.sh --after-remove ./uninstall.sh --deb-no-default-config-files /usr/lib/yeticloud/yeti-discover
fpm -t rpm -s dir -n yeti-discover -v "$1" -a amd64 -p "yeti-discover-$1-amd64.rpm" --license GPLv3 --vendor YetiCloud -m hello@yeticloud.com --url "https://github.com/yeticloud/yeti-discover" --description "YetiCloud yeti-discover binary distribution" --after-install ./install.sh --after-remove ./uninstall.sh /usr/lib/yeticloud/yeti-discover
