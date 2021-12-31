package network

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// IPTables fetches all iptables rules
func IPTables(d *data.DiscoverJSON) {

	if runtime.GOOS == "linux" {
		iptableL := exec.Command("iptables", "-L")
		iptableGrep := exec.Command("grep", "-v", "^Chain\\|target\\|^$")
		iptableLOut, err := iptableL.StdoutPipe()
		if err != nil {
			return
		}
		err = iptableL.Start()
		if err != nil {
			return
		}
		iptableGrep.Stdin = iptableLOut
		iptableOut, err := iptableGrep.Output()
		if err != nil {
			return
		}

		var iptableSlice []string

		for _, line := range strings.Split(strings.TrimSuffix(string(iptableOut), "\n"), "\n") {
			iptableSlice = append(iptableSlice, strings.TrimSpace(line))
		}

		d.Iptables = iptableSlice
	}
}
