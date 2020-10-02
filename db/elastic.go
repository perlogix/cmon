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

package db

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/yeticloud/yeti-discover/config"
	"github.com/yeticloud/yeti-discover/data"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Bool("verify_ssl")},
	}
	c = &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	host   = config.Str("host")
	port   = config.Str("port")
	env    = config.Str("environment")
	https  = config.Bool("https")
	user   = config.Str("username")
	pass   = config.Str("password")
	scheme string
	url    string
)

const ct = "application/json;charset=UTF-8"

func init() {
	if https {
		scheme = "https"
	} else {
		scheme = "http"
	}
}

// Elastic makes POST HTTP call to ElasticSearch DB
func Elastic(d *data.DiscoverJSON) {
	if url == "" {
		buf := new(bytes.Buffer)
		buf.WriteString(scheme)
		buf.WriteString("://")
		buf.WriteString(host)
		buf.WriteString(":")
		buf.WriteString(port)
		buf.WriteString("/servers/")
		buf.WriteString(env)
		url = buf.String()
	}

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(d)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	if user != "" && pass != "" {
		r, err := http.NewRequest("POST", url, body)
		if err != nil {
			log.Printf("Error: %s\n", err)
			return
		}
		r.Header.Add("Content-Type", ct)
		r.SetBasicAuth(user, pass)
		_, err = c.Do(r)
		if err != nil {
			log.Printf("Error: %s\n", err)
			return
		}
		return
	}

	_, err = c.Post(url, ct, body)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
