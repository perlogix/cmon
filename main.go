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

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yeticloud/yeti-discover/cloud"
	"github.com/yeticloud/yeti-discover/config"
	"github.com/yeticloud/yeti-discover/containers"
	"github.com/yeticloud/yeti-discover/data"
	"github.com/yeticloud/yeti-discover/db"
	"github.com/yeticloud/yeti-discover/network"
	"github.com/yeticloud/yeti-discover/packages"
	"github.com/yeticloud/yeti-discover/security"
	"github.com/yeticloud/yeti-discover/system"
)

var (
	configFlag = flag.String("config", "", "  Set configuration path, defaults are ['./', '/etc/yeticloud', '/opt/yeticloud', '/usr/lib/yeticloud/yeti-discover']")
	daemonFlag = flag.Bool("daemon", false, "  Run in daemon mode")
	builtOn    string
)

func init() {
	flag.StringVar(configFlag, "c", "", "  Set configuration path, defaults are ['./', '/etc/yeticloud', '/opt/yeticloud']")
	flag.BoolVar(daemonFlag, "d", false, "  Run in daemon mode")
}

func main() {

	flag.Usage = func() {
		fmt.Printf(` Usage: yeti-discover [options] <args>
   -d, --daemon    Run in daemon mode
   -c, --config    Set configuration path, defaults are ['./', '/etc/yeticloud', '/opt/yeticloud']
	
 Built On:       %s

 Example:        yeti-discover -d -c ./conf/yeti-discover.yaml
	
 Documentation:  https://github.com/yeticloud/yeti-discover/blob/master/README.md
`, builtOn)
	}

	flag.Parse()

	for {
		var (
			d  data.DiscoverJSON
			wg sync.WaitGroup
		)
		wg.Add(25)
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
			system.Stats(&d)
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
		// sysctl call
		go func() {
			defer wg.Done()
			system.Sysctl(&d)
		}()
		// lsmod call
		go func() {
			defer wg.Done()
			system.Lsmod(&d)
		}()
		wg.Wait()

		if !*daemonFlag {
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
