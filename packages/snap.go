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

package packages

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Snaps fetches all snap containers
func Snaps(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		snap := exec.Command("snap", "list")
		snapAwk := exec.Command("awk", "/^[a-z]/{print$1\"-\"$2}")
		snapOut, err := snap.StdoutPipe()
		if err != nil {
			return
		}
		err = snap.Start()
		if err != nil {
			return
		}
		snapAwk.Stdin = snapOut
		snapLOut, err := snapAwk.Output()
		if err != nil {
			return
		}

		snapSlice := strings.Split(strings.TrimSpace(string(snapLOut)), "\n")

		if snapSlice != nil {
			d.Snaps = snapSlice
		}
	}
}
