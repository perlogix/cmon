# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
    config.vm.box = "ubuntu/focal64"
    config.vm.hostname = "yeti-discover-stack"
    config.vm.network :private_network, ip: "10.20.1.32"
    config.vm.network "forwarded_port", guest: 5601, host: 5601, auto_correct: true
    config.vm.network "forwarded_port", guest: 9200, host: 9200, auto_correct: true
    config.vm.network "forwarded_port", guest: 9600, host: 9600, auto_correct: true
    config.vm.provider "virtualbox" do |v|
        v.name = "yeti-discover-stack"
        v.memory = 4096
        v.cpus = 2
        v.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
        v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        v.customize ["modifyvm", :id, "--uartmode1", "file", File::NULL]
    end
    config.vm.provision "shell", inline: <<-SHELL
# Setup
sysctl -w vm.max_map_count=262144
sysctl -p

apt-get update
apt-get install -y curl net-tools jq make ssg-debderived unzip wget apt-transport-https gnupg lsb-release clamav clamav-daemon

# Install Docker
curl -sSL https://get.docker.com/ |  sh
systemctl enable docker
systemctl restart docker

# Install Latest OpenSCAP Guides
wget -c https://github.com/ComplianceAsCode/content/releases/download/v0.1.53/scap-security-guide-0.1.53.zip
unzip scap-security-guide-0.1.53.zip
cp -rf scap-security-guide-0.1.53/* /usr/share/scap-security-guide/
rm -rf scap-security-guide-0.1.53*

# Install Trivy
wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list
apt-get update
apt-get install -y trivy

# Setup ClamAV
systemctl stop clamav-freshclam

# Update Virus Defs
freshclam
systemctl start clamav-freshclam
systemctl enable clamav-daemon clamav-freshclam

# Run OpenDistro ELK 
docker run -d --name es --net=host -e "discovery.type=single-node" -e "network.host=0.0.0.0" amazon/opendistro-for-elasticsearch:latest
docker run -d --name kibana --net=host amazon/opendistro-for-elasticsearch-kibana:latest
sleep 60

# Create Index servers
curl -XPUT -k -u admin:admin https://localhost:9200/servers

# Create ES Mapping for servers Index
curl -XPUT -k -u admin:admin "https://localhost:9200/servers/_mapping" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "asset_type": {
      "type": "keyword"
    },
    "chassis_type": {
      "type": "keyword"
    },
    "cpu_count": {
      "type": "long"
    },
    "cpu_pct": {
      "type": "long"
    },
    "clamav_defs": {
      "type": "keyword"
    },
    "crontabs": {
      "type": "keyword"
    },
    "diskfree_gb": {
      "type": "long"
    },
    "disktotal_gb": {
      "type": "long"
    },
    "diskused_gb": {
      "type": "long"
    },
    "dns_nameserver": {
      "type": "keyword"
    },
    "docker_containers": {
      "type": "keyword"
    },
    "docker_running": {
      "type": "long"
    },
    "docker_stopped": {
      "type": "long"
    },
    "docker_images_count": {
      "type": "long"
    },
    "docker_images": {
      "type": "keyword"
    },
    "environment": {
      "type": "keyword"
    },
    "expired_certs": {
      "type": "keyword"
    },
    "hostname": {
      "type": "keyword"
    },
    "ip_route": {
      "type": "keyword"
    },
    "ipaddress": {
      "type": "keyword"
    },
    "iptables": {
      "type": "keyword"
    },
    "network_interfaces": {
      "type": "nested",
      "properties": {
        "interface": {
          "type": "keyword"
        },
        "mtu": {
          "type": "long"
        },
        "rx_ok": {
          "type": "long"
        },
        "rx_err": {
          "type": "long"
        },
        "rx_drop": {
          "type": "long"
        },
        "rx_overrun": {
          "type": "long"
        },
        "tx_ok": {
          "type": "long"
        },
        "tx_err": {
          "type": "long"
        },
        "tx_drop": {
          "type": "long"
        },
        "tx_overrun": {
          "type": "long"
        },
        "flag": {
          "type": "keyword"
        }
      }
    },
    "kernel_version": {
      "type": "keyword"
    },
    "lastrun": {
      "type": "date"
    },
    "load1": {
      "type": "float"
    },
    "load15": {
      "type": "float"
    },
    "load5": {
      "type": "float"
    },
    "loaded_kernel_modules": {
      "type": "keyword"
    },
    "memoryfree_gb": {
      "type": "long"
    },
    "memorytotal_gb": {
      "type": "long"
    },
    "open_ports": {
      "type": "keyword"
    },
    "openscap": {
      "type": "object",
      "properties": {
        "status": {
          "type": "boolean"
        },
        "checks": {
          "type": "long"
        },
        "failed": {
          "type": "nested",
          "properties": {
            "title": {
              "type": "keyword"
            },
            "rule": {
              "type": "keyword"
            },
            "result": {
              "type": "keyword"
            }
          }
        },
        "warnings": {
          "type": "keyword"
        }
      }
    },
    "os": {
      "type": "keyword"
    },
    "packages": {
      "type": "keyword"
    },
    "platform": {
      "type": "keyword"
    },
    "platform_family": {
      "type": "keyword"
    },
    "platform_version": {
      "type": "float"
    },
    "processes": {
      "type": "keyword"
    },
    "public": {
      "type": "boolean"
    },
    "snaps": {
      "type": "keyword"
    },
    "sysctl": {
      "type": "keyword"
    },
    "systemd_timers": {
      "type": "keyword"
    },
    "trivy": {
      "type": "nested",
      "properties": {
        "Target": {
          "type": "keyword"
        },
        "Type": {
          "type": "keyword"
        },
        "Vulnerabilities": {
          "type": "nested",
          "properties": {
            "VulnerabilityID": {
              "type": "keyword"
            },
            "PkgName": {
              "type": "keyword"
            },
            "InstalledVersion": {
              "type": "keyword"
            },
            "Layer": {
              "type": "object",
              "properties": {
                "DiffID": {
                  "type": "keyword"
                }
              }
            },
            "SeveritySource": {
              "type": "keyword"
            },
            "PrimaryURL": {
              "type": "keyword"
            },
            "Title": {
              "type": "keyword"
            },
            "Description": {
              "type": "keyword"
            },
            "Severity": {
              "type": "keyword"
            },
            "CweIDs": {
              "type": "keyword"
            },
            "CVSS": {
              "type": "object",
              "properties": {
                "nvd": {
                  "type": "object",
                  "properties": {
                    "V2Vector": {
                      "type": "keyword"
                    },
                    "V3Vector": {
                      "type": "keyword"
                    },
                    "V2Score": {
                      "type": "float"
                    },
                    "V3Score": {
                      "type": "float"
                    }
                  }
                },
                "redhat": {
                  "type": "object",
                  "properties": {
                    "V3Vector": {
                      "type": "keyword"
                    },
                    "V3Score": {
                      "type": "float"
                    }
                  }
                }
              }
            },
            "References": {
              "type": "keyword"
            },
            "PublishedDate": {
              "type": "date"
            },
            "LastModifiedDate": {
              "type": "date"
            }
          }
        }
      }
    },
    "users": {
      "type": "keyword"
    },
    "users_loggedin": {
      "type": "keyword"
    },
    "virtualization": {
      "type": "boolean"
    },
    "virtualization_system": {
      "type": "keyword"
    }
  }
}
'

# Setup yeti-discover
mkdir -p /usr/lib/yeticloud/yeti-discover

cat <<'EOF'>/usr/lib/yeticloud/yeti-discover/yeti-discover.yml
host: 127.0.0.1
port: 9200
username: admin
password: admin
https: true
insecure_ssl: true
EOF

cp -f /vagrant/yeti-discover /usr/bin/

timeout 15 yeti-discover -d

# Test stack works, should return 2 
curl -sk -u admin:admin https://localhost:9200/servers/_search | jq -r '.hits.hits[0]._source.cpu_count'
SHELL
end