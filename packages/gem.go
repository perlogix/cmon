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
	"strings"

	"github.com/yeticloud/yeti-discover/data"
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
	gemStr := string(gemOut)
	gemReplace := strings.Replace(gemStr, " (", "-", -1)
	gemReplace2 := strings.Replace(gemReplace, ")", "", -1)
	gemSlice := strings.Split(strings.TrimSpace(gemReplace2), "\n")

	if gemSlice != nil {
		d.Gem = gemSlice
	}
}
