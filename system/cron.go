package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Cron fetches all crontabs
func Cron(d *data.DiscoverJSON) {
	findSlice := []string{}

	if runtime.GOOS == "linux" {

		cronOut, err := util.Cmd(`cat "/var/spool/cron/*/*" /etc/crontab /etc/anacrontab 2>/dev/null | grep -v '^#\|^[A-Z]\|^$'`)
		if err != nil {
			d.Crontabs = findSlice
			return
		}

		findSlice = append(findSlice, strings.Split(strings.Replace(strings.TrimSpace(string(cronOut)), "\t", " ", -1), "\n")...)
	}

	d.Crontabs = findSlice
}
