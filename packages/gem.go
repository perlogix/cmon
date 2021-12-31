package packages

import (
	"os/exec"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Gem fetches all gems on the system
func Gem(d *data.DiscoverJSON) {

	gemList := exec.Command("gem", "list")
	gemGrep := exec.Command("grep", "^[a-zA-Z]")
	gemListOut, err := gemList.StdoutPipe()
	if err != nil {
		return
	}
	err = gemList.Start()
	if err != nil {
		return
	}
	gemGrep.Stdin = gemListOut
	gemOut, err := gemGrep.Output()
	if err != nil {
		return
	}

	gemReplace := strings.Replace(string(gemOut), " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(strings.TrimSpace(gemReplace2), "\n")

	if gemSlice != nil {
		d.Gem = gemSlice
	}
}
