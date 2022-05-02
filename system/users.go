package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// UsersLoggedIn fetches all users logged into Linux or OS X
func UsersLoggedIn(d *data.DiscoverJSON) {
	if runtime.GOOS == "windows" {
		return
	}

	wOut, err := util.Cmd(`w -h | awk '{ print $1"-"$2}'`)
	if err != nil {
		return
	}

	wSlice := strings.Split(strings.TrimSpace(string(wOut)), "\n")

	d.UsersLoggedin = wSlice
}

// Users fetches all users on the system from /etc/passwd on NIX systems
func Users(d *data.DiscoverJSON) {
	if runtime.GOOS == "windows" {
		return
	}

	passwdOut, err := util.Cmd(`grep -v ^# /etc/passwd`)
	if err != nil {
		return
	}

	usersSlice := strings.Split(strings.TrimSpace(string(passwdOut)), "\n")
	if usersSlice != nil {
		d.Users = usersSlice
	}
}
