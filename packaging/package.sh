#!/bin/sh

if [ "$1" = "" ]; then
  echo "First argument empty, need version"
  exit 1
fi

mkdir -p /etc/perlogix/cmon

cp -f ./cmon.service /etc/perlogix/cmon/
cp -f ../cmon /etc/perlogix/cmon/

cat <<'EOF' >/etc/perlogix/cmon/cmon.yaml.orig
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

# Verify SSL for HTTPS endpoints
#insecure_ssl: false

# Public facing asset
#public: false

# Asset type
#asset_type:

# OpenSCAP XCCDF XML file path
#oscap_xccdf_xml: /usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml

# OpenSCAP Profile
#oscap_profile: xccdf_org.ssgproject.content_profile_cis
EOF

fpm -t deb -s dir -n cmon -v "$1" -a amd64 -p "cmon-$1-amd64.deb" --license GPLv3 --vendor Perlogix -m hello@perlogix.com --url "https://github.com/perlogix/cmon" --description "Perlogix cmon binary distribution" --after-install ./install.sh --after-remove ./uninstall.sh --deb-no-default-config-files /etc/perlogix/cmon
fpm -t rpm -s dir -n cmon -v "$1" -a amd64 -p "cmon-$1-amd64.rpm" --license GPLv3 --vendor Perlogix -m hello@perlogix.com --url "https://github.com/perlogix/cmon" --description "Perlogix cmon binary distribution" --after-install ./install.sh --after-remove ./uninstall.sh /etc/perlogix/cmon
