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

// Sysctl collects system sysctl kernel parameters
func Sysctl(d *data.DiscoverJSON) {
	if runtime.GOOS == "windows" {
		return
	}
	cmd := exec.Command("sysctl", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		return
	}
	outString := strings.Split(string(stdout), "\n")
	var outSlice []string
	var separator string
	if runtime.GOOS == "darwin" {
		separator = ":"
	} else if runtime.GOOS == "linux" {
		separator = "="
	}
	for _, s := range outString {
		if s != "" {
			key := strings.TrimSpace(strings.Split(s, separator)[0])
			value := strings.TrimSpace(strings.Split(s, separator)[1])
			outSlice = append(outSlice, key+"="+value)
		}
	}
	d.Sysctl = outSlice
}
