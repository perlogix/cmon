package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Cron fetches all crontabs
func Cron(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		cat := exec.Command("cat", "/var/spool/cron/*/*", "/etc/crontab", "/etc/anacrontab")
		catGrep := exec.Command("grep", "-v", "^#\\|^[A-Z]\\|^$")
		catGOut, err := cat.StdoutPipe()
		if err != nil {
			return
		}
		err = cat.Start()
		if err != nil {
			return
		}
		catGrep.Stdin = catGOut
		catOut, err := catGrep.Output()
		if err != nil {
			return
		}

		findSlice := strings.Split(strings.Replace(strings.TrimSpace(string(catOut)), "\t", " ", -1), "\n")

		if findSlice != nil {
			d.Crontabs = findSlice
		}
	}
}
