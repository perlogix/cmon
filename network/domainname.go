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

package network

import (
	"os"
	"strings"

	"github.com/perlogix/cmon/data"
)

// DomainName fetches the domain name used on system
func DomainName(d *data.DiscoverJSON) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	if strings.ContainsAny(hostname, ".") {
		d.Domain = strings.TrimSuffix(hostname, ".")
	}
}
