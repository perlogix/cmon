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

// TimeZone runs Linux command date to fetch timezone
func TimeZone(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		timezone, err := exec.Command("date", "+%Z").Output()
		if err != nil {
			return
		}

		timezoneTrim := strings.TrimSpace(string(timezone))

		if timezoneTrim != "" {
			d.Timezone = timezoneTrim
		}
	}
}
