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
	"github.com/perlogix/cmon/network"
	"github.com/perlogix/cmon/packages"
	"github.com/perlogix/cmon/security"
	"github.com/perlogix/cmon/shipper"
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
		wg.Add(43)
		go func() {
			defer wg.Done()
			system.Stats(&d)
		}()
		go func() {
			defer wg.Done()
			config.Tags(&d)
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
			network.NTPRunning(&d)
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
			packages.WindowsPackages(&d)
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
			security.CPUVulnerabilities(&d)
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
			system.DmesgErrors(&d)
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
		go func() {
			defer wg.Done()
			system.SystemctlFailed(&d)
		}()
		go func() {
			defer wg.Done()
			system.Journalctl(&d)
		}()
		go func() {
			defer wg.Done()
			d.ID = system.GetHostID()
		}()
		wg.Wait()

		if !config.Bool("daemon") {
			j, err := json.Marshal(d)
			if err != nil {
				log.Println("Error: ", err)
			}
			j = bytes.ReplaceAll(j, []byte("\\u003c"), []byte("<"))
			j = bytes.ReplaceAll(j, []byte("\\u003e"), []byte(">"))
			j = bytes.ReplaceAll(j, []byte("\\u0026"), []byte("&"))
			fmt.Println(string(j))
			return
		}

		shipper.Ship(&d)

		time.Sleep(time.Duration(config.Int("interval")) * time.Second)
	}
}
