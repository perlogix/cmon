package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// ClamAVDefs shows ClamAV version and signatures version followed by the date of the signatures
func ClamAVDefs(d *data.DiscoverJSON) {
	if runtime.GOOS == "linx" {
		out, err := exec.Command("clamscan", "--version").Output()
		if err != nil {
			return
		}

		d.ClamAVDefs = strings.TrimSuffix(string(out), "\n")
	}
}
