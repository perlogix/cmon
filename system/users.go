package system

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// UsersLoggedIn fetches all users logged into Linux or OS X
func UsersLoggedIn(d *data.DiscoverJSON) {
	wSlice := []string{}

	if runtime.GOOS == "windows" {
		d.UsersLoggedin = wSlice
		return
	}

	wOut, _ := util.Cmd(`w -h | awk '{ print $1"-"$2}'`)

	wSlice = append(wSlice, strings.Split(strings.TrimSpace(string(wOut)), "\n")...)

	d.UsersLoggedin = wSlice
}

// Users fetches all users on the system from /etc/passwd on NIX systems
func Users(d *data.DiscoverJSON) {
	usersSlice := []string{}

	if runtime.GOOS == "windows" {
		d.Users = usersSlice
		return
	}

	passwdOut, _ := util.Cmd(`grep -v ^# /etc/passwd`)

	usersSlice = append(usersSlice, strings.Split(strings.TrimSpace(string(passwdOut)), "\n")...)

	d.Users = usersSlice
}
