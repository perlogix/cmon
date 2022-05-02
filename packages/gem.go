package packages

import (
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Gem fetches all gems on the system
func Gem(d *data.DiscoverJSON) {

	gemOut, err := util.Cmd(`gem list | grep '^[a-zA-Z]'`)
	if err != nil {
		return
	}

	gemReplace := strings.Replace(string(gemOut), " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(strings.TrimSpace(gemReplace2), "\n")

	d.Gem = gemSlice
}
