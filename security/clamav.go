package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// ClamAVDefs shows ClamAV version and signatures version followed by the date of the signatures
func ClamAVDefs(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("clamscan", "--version").Output()
		if err != nil {
			return
		}

		d.ClamAVDefs = strings.TrimSuffix(string(out), "\n")
	}
}
