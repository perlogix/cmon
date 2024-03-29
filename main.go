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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/perlogix/cmon/cloud"
	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/containers"
	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/db"
	"github.com/perlogix/cmon/network"
	"github.com/perlogix/cmon/packages"
	"github.com/perlogix/cmon/security"
	"github.com/perlogix/cmon/system"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			main()
		}
	}()

	for {
		var (
			d  data.DiscoverJSON
			wg sync.WaitGroup
		)
		wg.Add(35)
		go func() {
			defer wg.Done()
			system.Stats(&d)
		}()
		go func() {
			defer wg.Done()
			network.Conns(&d)
		}()
		go func() {
			defer wg.Done()
			security.OScap(&d)
		}()
		go func() {
			defer wg.Done()
			network.IfaceStats(&d)
		}()
		go func() {
			defer wg.Done()
			security.ExpiredCerts(&d)
		}()
		go func() {
			defer wg.Done()
			security.FailedLogins(&d)
		}()
		go func() {
			defer wg.Done()
			network.NTPServers(&d)
		}()
		go func() {
			defer wg.Done()
			network.DNS(&d)
		}()
		go func() {
			defer wg.Done()
			network.DomainName(&d)
		}()
		go func() {
			defer wg.Done()
			network.IP(&d)
		}()
		go func() {
			defer wg.Done()
			network.IPRoutes(&d)
		}()
		go func() {
			defer wg.Done()
			network.IPTables(&d)
		}()
		go func() {
			defer wg.Done()
			cloud.AWS(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Deb(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Pip(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Pip3(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Rpm(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Gem(&d)
		}()
		go func() {
			defer wg.Done()
			packages.Snaps(&d)
		}()
		go func() {
			defer wg.Done()
			containers.DockerContainers(&d)
		}()
		go func() {
			defer wg.Done()
			security.ClamAVDefs(&d)
		}()
		go func() {
			defer wg.Done()
			containers.DockerServer(&d)
		}()
		go func() {
			defer wg.Done()
			containers.DockerImages(&d)
		}()
		go func() {
			defer wg.Done()
			system.Audit(&d)
		}()
		go func() {
			defer wg.Done()
			system.Users(&d)
		}()
		go func() {
			defer wg.Done()
			system.UsersLoggedIn(&d)
		}()
		go func() {
			defer wg.Done()
			security.TrivyScan(&d)
		}()
		go func() {
			defer wg.Done()
			system.Cron(&d)
		}()
		go func() {
			defer wg.Done()
			d.Lastrun = time.Now().Format(time.RFC3339)
		}()
		go func() {
			defer wg.Done()
			system.Sysctl(&d)
		}()
		go func() {
			defer wg.Done()
			system.Lsmod(&d)
		}()
		go func() {
			defer wg.Done()
			system.Processes(&d)
		}()
		go func() {
			defer wg.Done()
			system.ChassisType(&d)
		}()
		go func() {
			defer wg.Done()
			cloud.DetectCloud(&d)
		}()

		go func() {
			defer wg.Done()
			system.SystemdTimers(&d)
		}()
		wg.Wait()

		if !config.Bool("daemon") {
			j, err := json.Marshal(d)
			if err != nil {
				log.Printf("Error: %s\n", err)
			}
			j = bytes.ReplaceAll(j, []byte("\\u003c"), []byte("<"))
			j = bytes.ReplaceAll(j, []byte("\\u003e"), []byte(">"))
			j = bytes.ReplaceAll(j, []byte("\\u0026"), []byte("&"))
			fmt.Println(string(j))
			return
		}

		db.Elastic(&d)

		time.Sleep(time.Duration(config.Int("interval")) * time.Second)
	}
}
