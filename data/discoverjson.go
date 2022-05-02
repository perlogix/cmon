package data

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/types"
	docker "github.com/docker/docker/api/types"
)

// DiscoverJSON is the main struct for JSON Marshal
type DiscoverJSON struct {
	AssetType            string                 `json:"asset_type"`
	AuditRules           []string               `json:"audit_rules"`
	ChassisType          string                 `json:"chassis_type"`
	Cloud                string                 `json:"cloud"`
	CPUCount             int                    `json:"cpu_count"`
	CPUVulnerabilities   []string               `json:"cpu_vulnerabilities"`
	ClamAVDefs           string                 `json:"clamav_defs"`
	Crontabs             []string               `json:"crontabs"`
	Diskfree             uint64                 `json:"diskfree_gb"`
	Disktotal            uint64                 `json:"disktotal_gb"`
	Diskused             uint64                 `json:"diskused_gb"`
	Diskpct              int                    `json:"diskused_pct"`
	DNSNameserver        []string               `json:"dns_nameserver"`
	DmesgErrors          []string               `json:"dmesg_errors"`
	DockerContainers     []DockerContainersInfo `json:"docker_containers"`
	DockerRunning        int                    `json:"docker_running"`
	DockerPaused         int                    `json:"docker_paused"`
	DockerStopped        int                    `json:"docker_stopped"`
	DockerImagesCount    int                    `json:"docker_images_count"`
	DockerImages         []DockerImagesInfo     `json:"docker_images"`
	DockerLabels         []string               `json:"docker_labels"`
	Domain               string                 `json:"domain"`
	Ec2AmiID             string                 `json:"ec2_ami_id"`
	Ec2AvailabilityZone  string                 `json:"ec2_availability_zone"`
	Ec2InstanceID        string                 `json:"ec2_instance_id"`
	Ec2InstanceType      string                 `json:"ec2_instance_type"`
	Ec2Profile           string                 `json:"ec2_profile"`
	Ec2PublicIP4         string                 `json:"ec2_public_ip4"`
	Ec2SecurityGroups    []string               `json:"ec2_security_groups"`
	Environment          string                 `json:"environment"`
	ExpiredCerts         []string               `json:"expired_certs"`
	FailedLogins         []string               `json:"failed_logins"`
	Gem                  []string               `json:"gem"`
	Hostname             string                 `json:"hostname"`
	ID                   string                 `json:"id"`
	IPRoute              []string               `json:"ip_route"`
	Ipaddress            string                 `json:"ip_address"`
	Iptables             []string               `json:"iptables"`
	Interfaces           []IfaceData            `json:"network_interfaces"`
	Journalctl           []string               `json:"journalctl_logs"`
	KernelArch           string                 `json:"kernel_arch"`
	Kernelversion        string                 `json:"kernel_version"`
	Lastrun              string                 `json:"last_run"`
	Load1                float64                `json:"load1"`
	Load15               float64                `json:"load15"`
	Load5                float64                `json:"load5"`
	Lsmod                []string               `json:"loaded_kernel_modules"`
	Memoryfree           uint64                 `json:"memoryfree_gb"`
	Memorytotal          uint64                 `json:"memorytotal_gb"`
	Memoryused           uint64                 `json:"memoryused_gb"`
	Memoryusagepct       int                    `json:"memoryused_pct"`
	NTPServers           []string               `json:"ntp_servers"`
	NTPRunning           bool                   `json:"ntp_running"`
	OpenPorts            []OpenPorts            `json:"open_ports"`
	OpenScap             OScapOutput            `json:"openscap"`
	Os                   string                 `json:"os"`
	Packages             []string               `json:"packages"`
	Pip                  []string               `json:"pip"`
	Pip3                 []string               `json:"pip3"`
	Platform             string                 `json:"platform"`
	Platformfamily       string                 `json:"platform_family"`
	Platformversion      string                 `json:"platform_version"`
	Processes            []Processes            `json:"processes"`
	Public               bool                   `json:"public"`
	Snaps                []string               `json:"snaps"`
	Sysctl               []string               `json:"sysctl"`
	SystemctlFailed      []string               `json:"systemctl_failed"`
	SystemdTimers        []string               `json:"systemd_timers"`
	Tags                 []string               `json:"tags"`
	Timezone             string                 `json:"timezone"`
	Trivy                Trivy                  `json:"trivy"`
	Uptime               uint64                 `json:"uptime_days"`
	Users                []string               `json:"users"`
	UsersLoggedin        []string               `json:"users_loggedin"`
	WindowsPackages      []WindowsPackages      `json:"windows_software"`
	Virtualization       bool                   `json:"virtualization"`
	Virtualizationsystem string                 `json:"virtualization_system"`
}

// DockerContainersInfo contains container information
type DockerContainersInfo struct {
	Name    string        `json:"name"`
	Image   string        `json:"image"`
	Command string        `json:"command"`
	Ports   []docker.Port `json:"ports"`
	State   string        `json:"state"`
	Status  string        `json:"status"`
}

// DockerImagesInfo contains container image information
type DockerImagesInfo struct {
	Name    string `json:"name"`
	Size    string `json:"size"`
	Created string `json:"created"`
}

// OpenPorts contains open TCP/UDP port information
type OpenPorts struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

// OScapOutput is the output format for the parsed data
type OScapOutput struct {
	Status         bool          `json:"status"`
	Checks         int           `json:"checks"`
	PassTotal      int           `json:"pass_total"`
	FixedTotal     int           `json:"fixed_total"`
	InfoTotal      int           `json:"informational_total"`
	FailTotal      int           `json:"fail_total"`
	ErrorTotal     int           `json:"error_total"`
	UnknownTotal   int           `json:"unknown_total"`
	NotCheckTotal  int           `json:"notchecked_total"`
	NotSelectTotal int           `json:"notselected_total"`
	NotAppTotal    int           `json:"notapplicable_total"`
	Failed         []OScapResult `json:"failed"`
	Warnings       []string      `json:"warnings"`
}

// OScapResult holds the information about an individual check
type OScapResult struct {
	Title  string `json:"title"`
	Rule   string `json:"rule"`
	Result string `json:"result"`
}

// Trivy contains Trivy results
type Trivy struct {
	VulnToal     int           `json:"vulnerabilities_total"`
	VulnLow      int           `json:"vulnerabilities_low"`
	VulnMed      int           `json:"vulnerabilities_medium"`
	VulnHigh     int           `json:"vulnerabilities_high"`
	VulnCrit     int           `json:"vulnerabilities_critical"`
	VulnUnknown  int           `json:"vulnerabilities_unknown"`
	TrivyResults types.Results `json:"trivy_results"`
}

// IfaceData type is the Kernel Network Interface table
type IfaceData struct {
	Interface string `json:"interface"`
	MTU       int    `json:"mtu"`
	RXok      int    `json:"rx_ok"`
	RXerr     int    `json:"rx_err"`
	RXdrp     int    `json:"rx_drop"`
	RXovr     int    `json:"rx_overrun"`
	TXok      int    `json:"tx_ok"`
	TXerr     int    `json:"tx_err"`
	TXdrp     int    `json:"tx_drop"`
	TXovr     int    `json:"tx_overrun"`
	Flag      string `json:"flag"`
}

type Processes struct {
	Pid  int    `json:"pid"`
	Ppid int    `json:"ppid"`
	Name string `json:"name"`
	User string `json:"user"`
}

type WindowsPackages struct {
	DisplayName    string    `json:"display_name"`
	DisplayVersion string    `json:"display_version"`
	InstallDate    time.Time `json:"install_date"`
	Publisher      string    `json:"publisher"`
}
