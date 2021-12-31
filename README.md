# cmon

[![Go Report Card](https://goreportcard.com/badge/github.com/perlogix/cmon)](https://goreportcard.com/report/github.com/perlogix/cmon)
![Go](https://github.com/perlogix/cmon/workflows/Go/badge.svg)

#### Table of Contents

1. [Overview](#overview)
   - [Example JSON Output](#example-json-output)
2. [Install](#install)
3. [Install Dependencies](#install-dependencies)
   - [Server](#server)
   - [Client](#client-development)
4. [Getting Started Vagrant](#getting-started-vagrant)
5. [Command-Line Arguments](#command-line-arguments)
6. [Configuration](#configuration)
7. [Vagrant](#vagrant)
8. [Platforms Tested On](#platforms-tested-on)

## Overview

NIST Information Security Continuous Monitoring (ISCM) and configuration baseline data collector.

Great for keeping track of elastic environments, auditing or migrating servers by storing data in ElasticSearch or outputting to STDOUT.

Resources gathered if applicable:

- Asset Type
- Auditd Rules
- Chassis Type
- Cloud / Container Detection
- CPU Count
- CPU Stats
- CPU Vulnerabilities
- ClamAV Definitions
- Crontabs
- Disk Stats
- DMesg Errors
- Docker Containers
- Docker Images
- Docker Stats
- Domain Name
- EC2 Instance Metadata
- Environment
- Expired SSL Certs
- Failed Logins
- Ruby Gems
- Hostname
- IP Address
- IPTables Rules
- IP Routes
- Kernel Version
- Load Averages
- Loaded Kernel Modules
- Memory Stats
- Network Interface Stats
- NTP Servers
- NTP Running
- RPM / Deb Packages
- Python Pip Packages
- Public
- Snap Packages
- Sysctl Kernel Parameters
- Systemd Failed Services
- Systemd Timers
- Processes
- OpenSCAP XCCDF Scan
- OS Platform
- OS Family
- OS Version
- TCP 4/6 Listening Open Ports
- Timezone
- Trivy Scan
- Uptime
- Users
- Users Logged In
- Virtualization
- Virtualization System

### Example JSON Output

```json
{
  "audit_rules":[
    "-w /var/log/audit/ -p wa -k LOG_audit",
    "-w /etc/audit/auditd.conf -p wa -k CFG_audit",
    "-w /etc/rc.d/init.d/auditd -p wa -k CFG_audit",
    "-w /etc/sysconfig/auditd -p wa -k CFG_audit",
    "-w /etc/audit/audit.rules -p wa -k CFG_audit",
    "-w /etc/localtime -p wa -k time-change,CFG_system"
  ],
  "chassis_type":"notebook",
  "cloud":"k8s container",
  "cpu_count":4,
  "cpu_pct":76,
  "cpu_vulns":[
    "/sys/devices/system/cpu/vulnerabilities/spectre_v2:Vulnerable, IBPB: disabled, STIBP: disabled",
    "/sys/devices/system/cpu/vulnerabilities/itlb_multihit:KVM: Vulnerable"
  ],
  "clamav_defs": "ClamAV 0.102.4/26037/Sun Jan  3 12:37:28 2021",
  "crontabs":[
    "25 6 * * * root test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )",
    "47 6 * * 7 root test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.weekly )",
    "52 6 1 * * root test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.monthly )",
    "@monthly 15 cron.monthly run-parts --report /etc/cron.monthly"
  ],
  "diskfree_gb":6,
  "disktotal_gb":8,
  "diskused_gb":19,
  "dns_nameserver":[
    "8.8.8.8",
    "8.8.4.4"
  ],
  "dmesg_errors":"ACPI BIOS Error (bug): Failure creating named object [B.PCI0.RP17.PXSX.TBDU], AE_ALREADY_EXISTS (20200717/dswload2-326) ACPI Error: AE_ALREADY_EXISTS, During name lookup/catalog (20200717/psobject-220)",
  "docker_containers":[
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
  "domain":"ec2.internal",
  "ec2_ami_id":"ami-bc8131d4",
  "ec2_availability_zone":"us-east-1b",
  "ec2_instance_id":"i-1b8cc9cc",
  "ec2_instance_type":"t1.micro",
  "ec2_profile":"default-paravirtual",
  "ec2_public_ip4":"54.145.182.91",
  "ec2_security_groups":"default",
  "environment":"dev",
  "expired_certs":[
    "/etc/ssl/server.crt Certificate will expire",
    "/etc/nginx/server.crt Certificate will expire",
    "/etc/httpd/server.crt Certificate will expire"
  ],
  "failed_logins":[
    "root pts/1 Sun Jan  3 17:30 - 17:30  (00:00)"
  ],
  "gem":[
    "arr-pm-0.0.9",
    "backports-3.6.4",
    "cabin-0.7.1",
    "childprocess-0.5.6",
    "clamp-0.6.4",
    "ffi-1.9.8",
    "fpm-1.3.3",
    "json-1.8.2"
  ],
  "hostname":"ip-10-28-229-205",
  "ipaddress":"10.28.229.205",
  "iptables":[
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
  "ip_route":[
    "default via 192.168.1.1 dev eth0 ",
    "172.17.0.0/16 dev docker0  proto kernel  scope link  src 172.17.42.1 ",
    "192.168.1.0/24 dev eth0  proto kernel  scope link  src 192.168.1.10 "
  ],
  "network_interfaces":[
    {
      "interface":"enp60s0",
      "mtu":1500,
      "rx_ok":0,
      "rx_err":0,
      "rx_drop":0,
      "rx_overrun":0,
      "tx_ok":0,
      "tx_err":0,
      "tx_drop":0,
      "tx_overrun":0,
      "flag":"BMU"
    }
  ],
  "kernel_version":"2.6.32-431.29.2.el6.x86_64",
  "lastrun":"2015-05-21T23:29:49-04:00",
  "load15":0,
  "load1":0,
  "load5":0,
  "loaded_kernel_modules":[
    "uinput 20480 0",
    "binfmt_misc 16384 1"
  ],
  "memoryfree_gb":2,
  "memorytotal_gb":16,
  "memoryused_gb":14,
  "ntp_servers":[
    "server ntp.server.com"
  ],
  "ntp_running": true,
  "os":"linux",
  "packages":[
    "acl-2.2.49-6.el6.x86_64",
    "acpid-1.0.10-2.1.el6.x86_64",
    "alsa-lib-1.0.22-3.el6.x86_64",
    "atk-1.30.0-1.el6.x86_64"
  ],
  "pip":[
    "distribute-0.6.10",
    "Flask-0.10.1",
    "Flask-Limiter-0.7.4"
  ],
  "pip3":[
    "aiofiles-0.4.0",
    "aiohttp-3.3.2",
    "apturl-0.5.2"
  ],
  "platform":"centos",
  "platform_family":"rhel",
  "platform_verison":"6.5",
  "processes":[
    "pid=1 ppid=0 name=systemd user=root cpu_pct=0 mem_pct=0",
    "pid=2 ppid=0 name=kthreadd user=root cpu_pct=0 mem_pct=0"
  ],
  "public":false,
  "snaps":[
    "core-16-2.31.1",
    "slack-3.0.5"
  ],
  "sysctl":[
    "abi.vsyscall32=1",
    "debug.exception-trace=1"
  ],
  "systemctl_failed":[
    "0 loaded units listed."
  ],
  "systemd_timers":[
    "Wed 2020-12-09 17:31:09 EST 30min left Wed 2020-12-09 16:34:56 EST 25min ago anacron.timer anacron.service",
    "Wed 2020-12-09 19:56:18 EST 2h 56min left Wed 2020-12-09 13:57:56 EST 3h 2min ago fwupd-refresh.timer fwupd-refresh.service"
  ],
  "open_ports":[
    "addr=127.0.0.1 port=58494 name=code proto=tcp",
    "addr=0.0.0.0 port=5601 name=node proto=tcp",
    "addr=:: port=9200 name=0 proto=tcp"
  ],
  "openscap":{
    "status":false,
    "checks":71,
    "failed":[
      {
        "title":"Enable auditd Service",
        "rule":"xccdf_org.ssgproject.content_rule_service_auditd_enabled",
        "result":"fail"
      },
      {
        "title":"Ensure auditd Collects System Administrator Actions",
        "rule":"xccdf_org.ssgproject.content_rule_audit_rules_sysadmin_actions",
        "result":"fail"
      }
    ],
    "warnings":null
  },
  "timezone":"UTC",
  "trivy":[
    {
      "Target":"k3s (ubuntu 20.04)",
      "Type":"ubuntu",
      "Vulnerabilities":[
        {
          "VulnerabilityID":"CVE-2012-6655",
          "PkgName":"accountsservice",
          "InstalledVersion":"0.6.55-0ubuntu12~20.04.1",
          "Layer":{
            "DiffID":"sha256:7a32807fe5359af26e10053cb110e4a7576afa4a63c26d4af5ab763d6784fae7"
          },
          "SeveritySource":"ubuntu",
          "Title":"accountsservice: local encrypted password disclosure when changing password",
          "Description":"An issue exists AccountService 0.6.37 in the user_change_password_authorized_cb() function in user.c which could let a local users obtain encrypted passwords.",
          "Severity":"LOW",
          "CweIDs":[
            "CWE-732"
          ],
          "VendorVectors":{
            "nvd":{
              "v2":"AV:L/AC:L/Au:N/C:P/I:N/A:N",
              "v3":"CVSS:3.1/AV:L/AC:L/PR:L/UI:N/S:U/C:L/I:N/A:N"
            },
            "redhat":{
              "v2":"AV:L/AC:M/Au:N/C:P/I:N/A:N"
            }
          },
          "CVSS":{
            "nvd":{
              "V2Vector":"AV:L/AC:L/Au:N/C:P/I:N/A:N",
              "V3Vector":"CVSS:3.1/AV:L/AC:L/PR:L/UI:N/S:U/C:L/I:N/A:N",
              "V2Score":2.1,
              "V3Score":3.3
            },
            "redhat":{
              "V2Vector":"AV:L/AC:M/Au:N/C:P/I:N/A:N",
              "V2Score":1.9
            }
          },
          "References":[
            "http://openwall.com/lists/oss-security/2014/08/15/5",
            "http://www.openwall.com/lists/oss-security/2014/08/16/7",
            "http://www.securityfocus.com/bid/69245",
            "https://bugzilla.redhat.com/show_bug.cgi?id=CVE-2012-6655",
            "https://bugzilla.suse.com/show_bug.cgi?id=CVE-2012-6655",
            "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2012-6655",
            "https://exchange.xforce.ibmcloud.com/vulnerabilities/95325",
            "https://security-tracker.debian.org/tracker/CVE-2012-6655"
          ],
          "PublishedDate":"2019-11-27T18:15:00Z",
          "LastModifiedDate":"2019-12-16T19:47:00Z"
        }
      ]
    }
  ],
  "uptime_days":9,
  "users":[
    "root:x:0:0:root:/root:/bin/bash",
    "adm:x:3:4:adm:/var/adm:/sbin/nologin",
    "shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown",
    "nginx:x:998:997:Nginx web server:/var/lib/nginx:/sbin/nologin",
    "varnish:x:997:996:Varnish Cache:/var/lib/varnish:/sbin/nologin"
  ],
  "users_loggedin":[
    "root-pts/0",
    "timski-pts/1"
  ],
  "virtualization":true,
  "virtualization_system":"xen"
}
```

Average payload size: `200k`

ElasticSearch terminology:

http://elasticsearch:9200/index/type

Discover terminology:

http://elasticsearch:9200/servers/environment

Agent Run Time:

The agent runs every twenty minutes, and post real-time data to ElasticSearch.

If you were to delete all hosts in the environment nightly. If the agent is running and the server is up, it will populate the inventory currently with only running hosts and their data. This works very well in elastic compute environments.

Example with cURL:

If you want to manually / cron schedule cmon to post to ElasticSearch

```sh
# HTTP unauth
sudo ./cmon | curl -XPOST -H "Content-Type: application/json" -d @- "http://localhost:9200/servers/_doc/$(hostid)"


# Insecure SSL and basic auth
sudo ./cmon | curl -XPOST -k -u admin:admin -H "Content-Type: application/json" -d @- "https://localhost:9200/servers/_doc/$(hostid)"
```

## Install

Install the statically linked Linux binary:

```sh
curl -OL "https://github.com/perlogix/cmon/releases/download/1.1/cmon" && chmod -f 0755 ./cmon
```

Install DEB file:

```sh
curl -LO $(curl -s https://api.github.com/repos/perlogix/cmon/releases/latest | grep browser_download_url | grep deb | cut -d '"' -f 4)

dpkg -i ./cmon*.deb
```

Install RPM file:

```sh
curl -LO $(curl -s https://api.github.com/repos/perlogix/cmon/releases/latest | grep browser_download_url | grep rpm | cut -d '"' -f 4)

rpm -i ./cmon*.rpm
```

**ElasticSearch Mappings Needed**

```sh
# Create index
curl -XPUT "http://localhost:9200/servers"

# Put mappings to existing index
curl -XPUT "http://localhost:9200/servers/_mapping" -H 'Content-Type: application/json' -d@mapping.json
```

## Install Dependencies

### Server

- ElasticSearch 7.x
- Kibana 7.x

### Client (development)

- Go 1.15>=
- Make
- Docker (Optional)

To build the binary with Go run the following command:

```sh
make build
```

To build the binary with Docker run the following command:

```sh
make docker
```

To build the RPM and Deb packages with Docker run the following command:

```sh
make VER=1.1 pkgs
```

## Getting Started Vagrant

```sh
git clone https://github.com/perlogix/cmon.git

cd cmon

curl -LO "https://github.com/perlogix/cmon/releases/download/1.1/cmon"

vagrant up
```

## Command-Line Arguments

No flags / arguments will do a one-time run and produce a JSON file in the current path of the binary

    -d, --daemon     Run in daemon mode
    -c, --config     Set configuration path, defaults are ['./', '/opt/perlogix/cmon', '/etc/perlogix/cmon']

## Configuration

Configurations can be written in YAML, JSON or TOML.

_/etc/perlogix/cmon/cmon.yaml_
_DEFAULT values if no config is present_

```yaml
# ElasticSearch DB
host: localhost
port: 9200

# ElasticSearch Index Name
# This can be anything, it could be aws, softlayer, prod, staging
environment: dev

# Interval of agent runs in seconds
# Default is every twenty minutes
interval: 1200

# Username if http-basic plugin is enabled
username:

# Password if http-basic plugin is enabled
password:

# https true enables HTTPS instead of HTTP)
https: false

# Verify SSL for HTTPS endpoints
insecure_ssl: false

# Public facing asset
public: false

# Asset type
asset_type:

# OpenSCAP XCCDF XML file path
oscap_xccdf_xml: /usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml

# OpenSCAP Profile
oscap_profile: xccdf_org.ssgproject.content_profile_cis
```

## Vagrant

```sh
git clone https://github.com/perlogix/cmon.git
cd cmon
vagrant up
vagrant ssh
```

## Platforms Tested On

- CentOS/RHEL 7 - latest
- Fedora 20 - latest
- Ubuntu 16 - latest
- Mac OS X 16.7.0 - latest
- Windows 10 - latest
