package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/data"
)

// ChassisType detects the Chassis type like notebook, VM, server
func ChassisType(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}
	cmd := exec.Command("dmidecode", "--string", "chassis-type")
	stdout, err := cmd.Output()
	if err != nil {
		return
	}
	inString := strings.Split(string(stdout), "\n")
	var outString string

	for _, s := range inString {
		if s != "" {
			outString = strings.ToLower(strings.TrimSpace(s))
		}
	}

	assetType := config.Str("asset_type")
	if assetType == "" {
		d.AssetType = outString
	} else {
		d.AssetType = assetType
	}

	d.ChassisType = outString
}
