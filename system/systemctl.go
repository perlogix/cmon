package system

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// SystemctlFailed collects failed systemd units
func SystemctlFailed(d *data.DiscoverJSON) {
	outSlice := []string{}

	if runtime.GOOS == "linux" {

		systemctlOut, err := util.Cmd(`systemctl --failed --no-page | grep -v '^  UNIT\|listed'`)
		if err != nil {
			d.SystemctlFailed = outSlice
			return
		}

		outString := strings.Split(string(systemctlOut), "\n")

		for _, s := range outString {
			if s != "" {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(s, " ")
				outSlice = append(outSlice, s)
			}
		}
	}

	d.SystemctlFailed = outSlice
}
