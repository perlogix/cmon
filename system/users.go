// Copyright (C) YetiCloud
// This file is part of yeti-discover <https://github.com/yeticloud/yeti-discover>.
//
// yeti-discover is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// yeti-discover is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with yeti-discover.  If not, see <http://www.gnu.org/licenses/>.

package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
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
		whStr := string(wOut)
		wSlice := strings.Split(strings.TrimSpace(whStr), "\n")

		if wSlice != nil {
			d.UsersLoggedin = wSlice
		}
	}
}

// Users fetches all users on the system from /etc/passwd on NIX systems
func Users(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		passGrep := exec.Command("grep", "-v", "^#", "/etc/passwd")
		passGrepOut, err := passGrep.Output()
		if err != nil {
			return
		}
		passStr := string(passGrepOut)
		usersSlice := strings.Split(strings.TrimSpace(passStr), "\n")
		if usersSlice != nil {
			d.Users = usersSlice
		}
	}
}
