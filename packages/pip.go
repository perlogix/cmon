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

// Pip fetches all Python 2.x packages
func Pip(d *data.DiscoverJSON) {

	pipFree := exec.Command("pip", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		return
	}
	err = pipFree.Start()
	if err != nil {
		return
	}
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		return
	}
	pipStr := string(pipOut)
	pipReplace := strings.Replace(pipStr, "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if pipSlice != nil {
		d.Pip = pipSlice
	}

}

// Pip3 fetches all Python 3.x packages
func Pip3(d *data.DiscoverJSON) {

	pipFree := exec.Command("pip3", "freeze")
	pipSort := exec.Command("sort")
	pipFreeOut, err := pipFree.StdoutPipe()
	if err != nil {
		return
	}
	err = pipFree.Start()
	if err != nil {
		return
	}
	pipSort.Stdin = pipFreeOut
	pipOut, err := pipSort.Output()
	if err != nil {
		return
	}
	pipStr := string(pipOut)
	pipReplace := strings.Replace(pipStr, "==", "-", -1)
	pipSlice := strings.Split(strings.TrimSpace(pipReplace), "\n")

	if pipSlice != nil {
		d.Pip3 = pipSlice
	}

}
