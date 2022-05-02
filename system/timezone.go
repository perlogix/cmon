package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// TimeZone runs Linux command date to fetch timezone
func TimeZone(d *data.DiscoverJSON) {
	if runtime.GOOS == "windows" {
		return
	}

	dateOut, err := util.Cmd(`date +%Z`)
	if err != nil {
		return
	}

	timezoneTrim := strings.TrimSpace(string(dateOut))

	if timezoneTrim != "" {
		d.Timezone = timezoneTrim
	}
}
