package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// TimeZone runs Linux command date to fetch timezone
func TimeZone(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		timezone, err := exec.Command("date", "+%Z").Output()
		if err != nil {
			return
		}

		timezoneTrim := strings.TrimSpace(string(timezone))

		if timezoneTrim != "" {
			d.Timezone = timezoneTrim
		}
	}
}
