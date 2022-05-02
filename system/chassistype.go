package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// ChassisType detects the Chassis type like notebook, VM, server
func ChassisType(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {

		dmidecodeOut, err := util.Cmd(`dmidecode --string chassis-type`)
		if err != nil {
			return
		}

		inString := strings.Split(string(dmidecodeOut), "\n")
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
}
