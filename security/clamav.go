package security

import (
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// ClamAVDefs shows ClamAV version and signatures version followed by the date of the signatures
func ClamAVDefs(d *data.DiscoverJSON) {
	clamscan, err := util.Cmd(`clamscan --version`)
	if err != nil {
		return
	}

	d.ClamAVDefs = strings.TrimSuffix(string(clamscan), "\n")
}
