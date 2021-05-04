// Copyright (C) Perlogix
// This file is part of cmon <https://github.com/perlogix/cmon>.
//
// cmon is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cmon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cmon.  If not, see <http://www.gnu.org/licenses/>.

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
