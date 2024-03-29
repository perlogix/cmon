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
	"regexp"
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
)

// Lsmod lists the currently loaded kernel modules
func Lsmod(d *data.DiscoverJSON) {
	if runtime.GOOS != "linux" {
		return
	}
	c1 := exec.Command("lsmod")
	c2 := exec.Command("grep", "-v", "Module")
	stdout1, err := c1.StdoutPipe()
	if nil != err {
		return
	}
	err = c1.Start()
	if nil != err {
		return
	}
	c2.Stdin = stdout1
	stdout, err := c2.Output()
	if err != nil {
		return
	}
	var (
		outSlice  []string
		outString = strings.Split(string(stdout), "\n")
	)
	for _, s := range outString {
		if s != "" {
			space := regexp.MustCompile(`\s+`)
			s := space.ReplaceAllString(s, " ")
			outSlice = append(outSlice, s)
		}
	}
	d.Lsmod = outSlice
}
