package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// UsersLoggedIn fetches all users logged into Linux or OS X
func UsersLoggedIn(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		wh := exec.Command("w", "-h")
		whAwk := exec.Command("awk", "{print$1\"-\"$2}")
		whOut, err := wh.StdoutPipe()
		if err != nil {
			return
		}
		err = wh.Start()
		if err != nil {
			return
		}
		whAwk.Stdin = whOut
		wOut, err := whAwk.Output()
		if err != nil {
			return
		}

		wSlice := strings.Split(strings.TrimSpace(string(wOut)), "\n")

		if wSlice != nil {
			d.UsersLoggedin = wSlice
		}
	}
}

// Users fetches all users on the system from /etc/passwd on NIX systems
func Users(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		passGrep, err := exec.Command("grep", "-v", "^#", "/etc/passwd").Output()
		if err != nil {
			return
		}

		usersSlice := strings.Split(strings.TrimSpace(string(passGrep)), "\n")
		if usersSlice != nil {
			d.Users = usersSlice
		}
	}
}
