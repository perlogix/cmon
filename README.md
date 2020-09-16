# yeti-discover

[![Go Report Card](https://goreportcard.com/badge/github.com/yeticloud/yeti-discover)](https://goreportcard.com/report/github.com/yeticloud/yeti-discover)

#### Table of Contents

1. [Overview](#overview)
	* [Example JSON Output](#example-json-output)
2. [Install](#install)
3. [Install Dependencies](#install-dependencies)
    * [Server](#server)
    * [Client](#client)
4. [Command-Line Arguments](#command-line-arguments)
5. [Configuration](#configuration)
6. [Platforms Tested On](#platforms-tested-on)
7. [Screenshots](#screenshots)


## Overview

A lightweight system information collector for storing data in ElasticSearch or Stdout.  Great for keeping track of elastic environments and auditing configurations.

Resources gathered if applicable:

- RHEL Audit Rules
- CPU Count
- CPU Stats
- Disk Stats
- Docker Containers
- Docker Images
- Docker Stats
- Domain Name
- EC2 Instance Metadata
- Ruby Gems
- Hostname
- IP Address
- IPTables Rules
- IP Routes
- Kernel Version
- Load Averages
- Memory Stats
- RPM / Deb Packages
- Python Pip Packages
- Snap Packages
- OpenSCAP XCCDF Scan
- OS Platform
- OS Family
- OS Version
- TCP 4/6 Listening
- Timezone
- Uptime
- Users
- Users Logged In
- Virtualization
- Virtualization System


### Example JSON Output

    {
       "audit_rules": [
         "-w /var/log/audit/ -p wa -k LOG_audit",
         "-w /etc/audit/auditd.conf -p wa -k CFG_audit",
         "-w /etc/rc.d/init.d/auditd -p wa -k CFG_audit",
         "-w /etc/sysconfig/auditd -p wa -k CFG_audit",
         "-w /etc/audit/audit.rules -p wa -k CFG_audit",
         "-w /etc/localtime -p wa -k time-change,CFG_system"
       ],
       "cpu_count": 4,
       "cpu_pct": 76,
       "diskfree_gb": 6,
       "disktotal_gb": 8,
       "diskused_gb": 19,
       "dns_nameserver": [
         "8.8.8.8",
         "8.8.4.4"
       ],
       "docker_containers": [
        "name=kibana image=kibana:7.4.0 command=/usr/local/bin/dumb-init -- /usr/local/bin/kibana-docker ports=[] state=running status=Up About a minute",
        "name=elasticsearch image=elasticsearch:7.4.0 command=/usr/local/bin/docker-entrypoint.sh eswrapper ports=[] state=running status=Up 3 minutes",
        "name=redis image=redis command=docker-entrypoint.sh redis-server ports=[{127.0.0.1 6379 6379 tcp}] state=running status=Up About an hour"
       ],
       "docker_running":3,
       "docker_images_count":3,
       "docker_images":[
        "name=kibana:7.4.0 size=1.097GB created=2019-09-27T05:25:49-04:00",
        "name=elasticsearch:7.4.0 size=858.7MB created=2019-09-27T04:42:16-04:00",
        "name=redis:latest size=95MB created=2019-03-26T20:49:00-04:00"
       ],
       "domain": "ec2.internal",
       "ec2_ami_id": "ami-bc8131d4",
       "ec2_availability_zone": "us-east-1b",
       "ec2_instance_id": "i-1b8cc9cc",
       "ec2_instance_type": "t1.micro",
       "ec2_profile": "default-paravirtual",
       "ec2_public_ip4": "54.145.182.91",
       "ec2_security_groups": "default",
       "environment": "dev",
       "gem": [
         "arr-pm-0.0.9",
         "backports-3.6.4",
         "cabin-0.7.1",
         "childprocess-0.5.6",
         "clamp-0.6.4",
         "ffi-1.9.8",
         "fpm-1.3.3",
         "json-1.8.2"
       ],
       "hostname": "ip-10-28-229-205",
       "ipaddress": "10.28.229.205",
       "iptables": [
         "ACCEPT     tcp  --  anywhere             anywhere             state RELATED,ESTABLISHED",
         "DROP       all  -f  anywhere             anywhere            ",
         "ACCEPT     tcp  --  localhost            anywhere             tcp dpt:webcache",
         "ACCEPT     tcp  --  localhost            anywhere             tcp dpt:webcache",
         "DROP       tcp  --  anywhere             anywhere             tcp dpt:webcache",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp dpt:http state NEW,ESTABLISHED",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp dpt:http limit: avg 25/min burst 100",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp spt:http state ESTABLISHED",
         "ACCEPT     tcp  --  anywhere             anywhere             tcp spt:webcache state ESTABLISHED"
       ],
       "ip_route": [
         "default via 192.168.1.1 dev eth0 ",
         "172.17.0.0/16 dev docker0  proto kernel  scope link  src 172.17.42.1 ",
         "192.168.1.0/24 dev eth0  proto kernel  scope link  src 192.168.1.10 "
       ],
       "kernel_version": "2.6.32-431.29.2.el6.x86_64",
       "lastrun": "2015-05-21T23:29:49-04:00",
       "load15": 0,
       "load1": 0,
       "load5": 0,
       "memoryfree_gb": 2,
       "memorytotal_gb": 16,
       "memoryused_gb": 14,
       "os": "linux",
       "packages": [
         "acl-2.2.49-6.el6.x86_64",
         "acpid-1.0.10-2.1.el6.x86_64",
         "alsa-lib-1.0.22-3.el6.x86_64",
         "atk-1.30.0-1.el6.x86_64"
       ],
       "pip": [
         "distribute-0.6.10",
         "Flask-0.10.1",
         "Flask-Limiter-0.7.4"
       ],
       "pip3":[
        "aiofiles-0.4.0",
        "aiohttp-3.3.2",
        "apturl-0.5.2"
       ]  
       "platform": "centos",
       "platform_family": "rhel",
       "platform_verison": "6.5",
       "public": false,
       "snaps":[
         "core-16-2.31.1",
         "slack-3.0.5",
       ],
       "open_ports":[
        "addr=127.0.0.1 port=58494 name=code proto=tcp",
        "addr=0.0.0.0 port=5601 name=node proto=tcp",
        "addr=:: port=9200 name=0 proto=tcp",
       ],
      "openscap": {
        "status": false,
        "checks": 71,
        "failed": [
          {
            "title": "Enable auditd Service",
            "rule": "xccdf_org.ssgproject.content_rule_service_auditd_enabled",
            "result": "fail"
          },
          {
            "title": "Ensure auditd Collects System Administrator Actions",
            "rule": "xccdf_org.ssgproject.content_rule_audit_rules_sysadmin_actions",
            "result": "fail"
          }
        ],
        "warnings": null
      },
       "timezone": "UTC",
       "uptime_days": 9,
       "users": [
         "root:x:0:0:root:/root:/bin/bash",
         "adm:x:3:4:adm:/var/adm:/sbin/nologin",
         "shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown",
         "nginx:x:998:997:Nginx web server:/var/lib/nginx:/sbin/nologin",
         "varnish:x:997:996:Varnish Cache:/var/lib/varnish:/sbin/nologin"
       ],
       "users_loggedin": [
         "root-pts/0",
         "timski-pts/1"
       ],
       "virtualization": true,
       "virtualization_system": "xen"
    }


ElasticSearch terminology:

http://elasticsearch:9200/index/type

 Discover terminology:

http://elasticsearch:9200/servers/environment

Agent Run Time:

The agent runs every five minutes, and post real-time data to ElasticSearch.

** If you were to delete all hosts in the environment nightly.   If the agent is running and the server is up, it will populate the inventory currently with only running hosts and their data.  This works very well in elastic compute environments.


## Install

Install the statically linked Linux binary:

    curl -OL https://github.com/yeticloud/yeti-discover/releases/download/0.1/yeti-discover && chmod -f 0755 ./yeti-discover


## Install Dependencies

### Server

 - ElasticSearch 7.x
 - Kibana 7.x

### Client (development)

 - Go 1.15.x
 - Make
 
To build the Linux binary run the following command:

    make linux

## Command-Line Arguments

No flags / arguments will do a one-time run and produce a JSON file in the current path of the binary

    -d, --daemon     Run in daemon mode
    -c, --config     Set configuration path, defaults are ['./','/etc/yeticloud','/opt/yeticloud']


## Configuration

Configurations can be written in YAML, JSON or TOML.

*/opt/yeticloud/yeti-discover.yaml*
*DEFAULT  values if no config is present*

    # ElasticSearch DB
    host: localhost
    port: 9200

    # ElasticSearch Index Name
    ## This can be anything, it could be aws, softlayer, prod, staging
    environment: dev

    # Interval of agent runs in seconds
    ## Default is every five minutes
    interval: 300

    # Username if http-basic plugin is enabled
    username:

    # Password if http-basic plugin is enabled
    password:

    # Secure true enables HTTPS instead of HTTP)
    secure: false

    # Verify SSL for HTTP endpoints
    verify_ssl: true

    # Public facing asset
    public: false

    # Asset type
    asset_type:

    # OpenSCAP XCCDF XML file path
    oscap_xccdf_xml: /usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml

    # OpenSCAP Profile
    oscap_profile: xccdf_org.ssgproject.content_profile_cis


## Platforms Tested On

 - CentOS/RHEL 7-latest
 - Fedora 20-latest
 - Ubuntu 16-latest
 - Mac OS X 16.7.0-latest

## Screenshots
![First View](https://yeticloud-public.s3.amazonaws.com/yeti-discover-kibana-1.png)

![Second View](https://yeticloud-public.s3.amazonaws.com/yeti-discover-kibana-2.png)

![Third View](https://yeticloud-public.s3.amazonaws.com/yeti-discover-kibana-3.png)
