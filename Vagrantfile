# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
    config.vm.box = "ubuntu/focal64"
    config.vm.hostname = "cmon-stack"
    config.vm.network :private_network, ip: "10.20.1.32"
    config.vm.network "forwarded_port", guest: 5601, host: 5601, auto_correct: true
    config.vm.network "forwarded_port", guest: 9200, host: 9200, auto_correct: true
    config.vm.network "forwarded_port", guest: 9600, host: 9600, auto_correct: true
    config.vm.provider "virtualbox" do |v|
        v.name = "cmon-stack"
        v.memory = 4096
        v.cpus = 2
        v.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
        v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        v.customize ["modifyvm", :id, "--uartmode1", "file", File::NULL]
    end
    config.vm.provision "shell", inline: <<-SHELL
# Setup
apt-get update -y
apt-get install -y curl
curl -LO https://raw.githubusercontent.com/perlogix/opsy/main/linux-ops/install.sh
chmod 0755 ./install.sh
./install.sh

apt-get install -y make ssg-debderived unzip wget apt-transport-https gnupg lsb-release clamav clamav-daemon

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
docker run -d --restart=always --name es --net=host -e "discovery.type=single-node" -e "network.host=0.0.0.0" amazon/opendistro-for-elasticsearch:latest
docker run -d --restart=always --name kibana --net=host amazon/opendistro-for-elasticsearch-kibana:latest
sleep 60

# Create Index servers
curl -XPUT -k -u admin:admin https://localhost:9200/servers

# Create ES Mapping for servers Index
curl -XPUT -k -u admin:admin "https://localhost:9200/servers/_mapping" -H 'Content-Type: application/json' -d@/vagrant/mapping.json

# Setup cmon
if [[ -f "/vagrant/cmon" ]]; then
  mkdir -p /etc/perlogix/cmon

cat <<'EOF'>/etc/perlogix/cmon/cmon.yaml
host: 127.0.0.1
port: 9200
username: admin
password: admin
https: true
insecure_ssl: true
EOF

  cp -f /vagrant/cmon /usr/bin/
else
  curl -LO $(curl -s https://api.github.com/repos/perlogix/cmon/releases/latest | grep browser_download_url | grep deb | cut -d '"' -f 4)
  dpkg -i ./cmon*.deb
fi

cp -f /vagrant/cmon /usr/bin/

timeout 15 cmon -d

# Test stack works, should return 2 
curl -sk -u admin:admin https://localhost:9200/servers/_search | jq -r '.hits.hits[0]._source.cpu_count'
SHELL
end