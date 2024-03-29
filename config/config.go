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

package config

import (
	"flag"
	"fmt"

	"github.com/denisbrodbeck/machineid"
	"github.com/spf13/viper"
)

var (
	v          = viper.New()
	configFlag = flag.String("config", "", "  Set configuration path, defaults are ['./', '/opt/perlogix/cmon', '/etc/perlogix/cmon']")
	daemonFlag = flag.Bool("daemon", false, "  Run in daemon mode")
	builtOn    string
	commitHash string
	version    string
	id, _      = machineid.ProtectedID("cmon")
)

func init() {

	flag.StringVar(configFlag, "c", "", "  Set configuration path, defaults are ['./', '/opt/perlogix/cmon', '/etc/perlogix/cmon']")
	flag.BoolVar(daemonFlag, "d", false, "  Run in daemon mode")

	flag.Usage = func() {
		fmt.Printf(` Usage: cmon [options] <args>
   -d, --daemon    Run in daemon mode
   -c, --config    Set configuration path, defaults are ['./', '/opt/perlogix/cmon', '/etc/perlogix/cmon']

 Version:        %s
 Built On:       %s
 Commit Hash:    %s

 Example:        cmon -d -c ./cmon.yaml
	
 Documentation:  https://github.com/perlogix/cmon/blob/master/README.md
`, version, builtOn, commitHash)
	}

	flag.Parse()

	// Default settings if no config file is supplied
	v.SetDefault("daemon", *daemonFlag)
	v.SetDefault("host", "localhost")
	v.SetDefault("port", "9200")
	v.SetDefault("environment", "dev")
	v.SetDefault("interval", "1200")
	v.SetDefault("username", "")
	v.SetDefault("password", "")
	v.SetDefault("https", "false")
	v.SetDefault("insecure_ssl", "false")
	v.SetDefault("public", "false")
	v.SetDefault("asset_type", "")
	v.SetDefault("scheme", "http")
	v.SetDefault("hostid", id)
	v.SetDefault("oscap_profile", "xccdf_org.ssgproject.content_profile_cis")
	v.SetDefault("oscap_xccdf_xml", "/usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml")

	v.SetConfigName("cmon")
	v.AddConfigPath("/etc/perlogix/cmon")
	v.AddConfigPath("/opt/perlogix/cmon")
	v.AddConfigPath(".")
	v.SetConfigFile(*configFlag)

	// Ignore error
	_ = v.ReadInConfig()
}

// Str fetches String value from configuration key
func Str(key string) string {
	return v.GetString(key)
}

// Int fetches Int value from configuration key
func Int(key string) int {
	return v.GetInt(key)
}

// Bool fetches Boolean value from configuration key
func Bool(key string) bool {
	return v.GetBool(key)
}
