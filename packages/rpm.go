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

package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// Rpm fetches all RPM packages installed on the system
func Rpm(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		rpmqa := exec.Command("rpm", "-qa")
		rpmSort := exec.Command("sort")
		rpmqaOut, err := rpmqa.StdoutPipe()
		if err != nil {
			return
		}
		err = rpmqa.Start()
		if err != nil {
			return
		}
		rpmSort.Stdin = rpmqaOut
		rpmOut, err := rpmSort.Output()
		if err != nil {
			return
		}

		rpmSlice := strings.Split(strings.TrimSpace(string(rpmOut)), "\n")

		if rpmSlice != nil {
			d.Packages = rpmSlice
		}
	}
}
