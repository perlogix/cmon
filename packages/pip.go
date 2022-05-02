package packages

import (
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// Pip fetches all Python 2.x packages
func Pip(d *data.DiscoverJSON) {

	pipOut, err := util.Cmd(`pip freeze | sort`)
	if err != nil {
		return
	}

	pipReplace := strings.Replace(string(pipOut), "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	d.Pip = pipSlice

}

// Pip3 fetches all Python 3.x packages
func Pip3(d *data.DiscoverJSON) {

	pipOut, err := util.Cmd(`pip3 freeze | sort`)
	if err != nil {
		return
	}

	pipReplace := strings.Replace(string(pipOut), "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	d.Pip3 = pipSlice

}
