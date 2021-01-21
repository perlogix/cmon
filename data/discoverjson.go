// Copyright (C) YetiCloud
// This file is part of yeti-discover <https://github.com/yeticloud/yeti-discover>.
//
// yeti-discover is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// yeti-discover is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with yeti-discover.  If not, see <http://www.gnu.org/licenses/>.

package data

import (
	ftypes "github.com/aquasecurity/fanal/types"
	"github.com/aquasecurity/trivy/pkg/types"
)

// DiscoverJSON is the main struct for JSON Marshal
type DiscoverJSON struct {
	AssetType            string      `json:"asset_type,omitempty"`
	AuditRules           []string    `json:"audit_rules,omitempty"`
	ChassisType          string      `json:"chassis_type,omitempty"`
	Cloud                string      `json:"cloud,omitempty"`
	CPUCount             int         `json:"cpu_count,omitempty"`
	CPUPct               int         `json:"cpu_pct,omitempty"`
	CPUvulns             []string    `json:"cpu_vulns,omitempty"`
	ClamAVDefs           string      `json:"clamav_defs,omitempty"`
	Crontabs             []string    `json:"crontabs,omitempty"`
	Diskfree             uint64      `json:"diskfree_gb,omitempty"`
	Disktotal            uint64      `json:"disktotal_gb,omitempty"`
	Diskused             uint64      `json:"diskused_gb,omitempty"`
	DNSNameserver        []string    `json:"dns_nameserver,omitempty"`
	DmesgErrors          string      `json:"dmesg_errors,omitempty"`
	DockerContainers     []string    `json:"docker_containers,omitempty"`
	DockerRunning        int         `json:"docker_running,omitempty"`
	DockerPaused         int         `json:"docker_paused,omitempty"`
	DockerStopped        int         `json:"docker_stopped,omitempty"`
	DockerImagesCount    int         `json:"docker_images_count,omitempty"`
	DockerImages         []string    `json:"docker_images,omitempty"`
	DockerLabels         []string    `json:"docker_labels,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Ec2AmiID             string      `json:"ec2_ami_id,omitempty"`
	Ec2AvailabilityZone  string      `json:"ec2_availability_zone,omitempty"`
	Ec2InstanceID        string      `json:"ec2_instance_id,omitempty"`
	Ec2InstanceType      string      `json:"ec2_instance_type,omitempty"`
	Ec2Profile           string      `json:"ec2_profile,omitempty"`
	Ec2PublicIP4         string      `json:"ec2_public_ip4,omitempty"`
	Ec2SecurityGroups    []string    `json:"ec2_security_groups,omitempty"`
	Environment          string      `json:"environment,omitempty"`
	ExpiredCerts         []string    `json:"expired_certs,omitempty"`
	FailedLogins         []string    `json:"failed_logins,omitempty"`
	Gem                  []string    `json:"gem,omitempty"`
	Hostname             string      `json:"hostname,omitempty"`
	IPRoute              []string    `json:"ip_route,omitempty"`
	Ipaddress            string      `json:"ipaddress,omitempty"`
	Iptables             []string    `json:"iptables,omitempty"`
	Interfaces           []IfaceData `json:"network_interfaces,omitempty"`
	Kernelversion        string      `json:"kernel_version,omitempty"`
	Lastrun              string      `json:"lastrun,omitempty"`
	Load1                float64     `json:"load1,omitempty"`
	Load15               float64     `json:"load15,omitempty"`
	Load5                float64     `json:"load5,omitempty"`
	Lsmod                []string    `json:"loaded_kernel_modules,omitempty"`
	Memoryfree           uint64      `json:"memoryfree_gb,omitempty"`
	Memorytotal          uint64      `json:"memorytotal_gb,omitempty"`
	Memoryused           uint64      `json:"memoryused_gb,omitempty"`
	NTPServers           []string    `json:"ntp_servers,omitempty"`
	NTPRunning           bool        `json:"ntp_running,omitempty"`
	OpenPorts            []string    `json:"open_ports,omitempty"`
	OpenScap             OScapOutput `json:"openscap,omitempty"`
	Os                   string      `json:"os,omitempty"`
	Packages             []string    `json:"packages,omitempty"`
	Pip                  []string    `json:"pip,omitempty"`
	Pip3                 []string    `json:"pip3,omitempty"`
	Platform             string      `json:"platform,omitempty"`
	Platformfamily       string      `json:"platform_family,omitempty"`
	Platformversion      string      `json:"platform_version,omitempty"`
	Processes            []string    `json:"processes,omitempty"`
	Public               bool        `json:"public,omitempty"`
	Snaps                []string    `json:"snaps,omitempty"`
	Sysctl               []string    `json:"sysctl,omitempty"`
	SystemctlFailed      string      `json:"systemctl_failed,omitempty"`
	SystemdTimers        []string    `json:"systemd_timers,omitempty"`
	Timezone             string      `json:"timezone,omitempty"`
	Trivy                Trivy       `json:"trivy,omitempty"`
	Uptime               uint64      `json:"uptime_days,omitempty"`
	Users                []string    `json:"users,omitempty"`
	UsersLoggedin        []string    `json:"users_loggedin,omitempty"`
	Virtualization       bool        `json:"virtualization,omitempty"`
	Virtualizationsystem string      `json:"virtualization_system,omitempty"`
}

// OScapOutput is the output format for the parsed data
type OScapOutput struct {
	Status         bool          `json:"status,omitempty"`
	Checks         int           `json:"checks,omitempty"`
	PassTotal      int           `json:"pass_total,omitempty"`
	FixedTotal     int           `json:"fixed_total,omitempty"`
	InfoTotal      int           `json:"informational_total,omitempty"`
	FailTotal      int           `json:"fail_total,omitempty"`
	ErrorTotal     int           `json:"error_total,omitempty"`
	UnknownTotal   int           `json:"unknown_total,omitempty"`
	NotCheckTotal  int           `json:"notchecked_total,omitempty"`
	NotSelectTotal int           `json:"notselected_total,omitempty"`
	NotAppTotal    int           `json:"notapplicable_total,omitempty"`
	Failed         []OScapResult `json:"failed,omitempty"`
	Warnings       []string      `json:"warnings,omitempty"`
}

// OScapResult holds the information about an individual check
type OScapResult struct {
	Title  string `json:"title,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Result string `json:"result,omitempty"`
}

// Trivy contains Trivy results
type Trivy struct {
	VulnToal     int `json:"vulnerabilities_total,omitempty"`
	VulnLow      int `json:"vulnerabilities_low,omitempty"`
	VulnMed      int `json:"vulnerabilities_medium,omitempty"`
	VulnHigh     int `json:"vulnerabilities_high,omitempty"`
	VulnCrit     int `json:"vulnerabilities_critical,omitempty"`
	VulnUnknown  int `json:"vulnerabilities_unknown,omitempty"`
	TrivyResults []TrivyResult
}

// TrivyResult type imported from Trivy go project
type TrivyResult struct {
	Target          string                        `json:"Target,omitempty"`
	Type            string                        `json:"Type,omitempty"`
	Packages        []ftypes.Package              `json:"Packages,omitempty"`
	Vulnerabilities []types.DetectedVulnerability `json:"Vulnerabilities,omitempty"`
}

// IfaceData type is the Kernel Network Interface table
type IfaceData struct {
	Interface string `json:"interface,omitempty"`
	MTU       int    `json:"mtu,omitempty"`
	RXok      int    `json:"rx_ok,omitempty"`
	RXerr     int    `json:"rx_err,omitempty"`
	RXdrp     int    `json:"rx_drop,omitempty"`
	RXovr     int    `json:"rx_overrun,omitempty"`
	TXok      int    `json:"tx_ok,omitempty"`
	TXerr     int    `json:"tx_err,omitempty"`
	TXdrp     int    `json:"tx_drop,omitempty"`
	TXovr     int    `json:"tx_overrun,omitempty"`
	Flag      string `json:"flag,omitempty"`
}
