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

// Audit fetches audit rules from auditctl -l Linux command
func Audit(d *data.DiscoverJSON) {
	if runtime.GOOS == "linux" {
		audit := exec.Command("auditctl", "-l")
		auditOut, err := audit.Output()
		if err != nil {
			return
		}
		auditStr := string(auditOut)
		auditSlice := strings.Split(strings.TrimSpace(auditStr), "\n")

		if auditSlice != nil {
			d.AuditRules = auditSlice
		}
	}
}
