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

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	// Configuration file settings using key-value
	viper.SetConfigName("yeti-discover")
	viper.AddConfigPath("/opt/yeticloud")
	viper.AddConfigPath("/etc/yeticloud")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println()
	}

	// Default settings if no config file is supplied
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "9200")
	viper.SetDefault("environment", "dev")
	viper.SetDefault("interval", "300")
	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.SetDefault("secure", "false")
	viper.SetDefault("verify_ssl", "true")
	viper.SetDefault("oscap_profile", "xccdf_org.ssgproject.content_profile_cis")
	viper.SetDefault("oscap_xccdf_xml", "/usr/share/scap-security-guide/ssg-ubuntu1804-ds.xml")
}

// Str fetches String value from configuration key
func Str(key string) string {
	return viper.GetString(key)
}

// Int fetches Int value from configuration key
func Int(key string) int {
	return viper.GetInt(key)
}

// Bool fetches Boolean value from configuration key
func Bool(key string) bool {
	return viper.GetBool(key)
}
