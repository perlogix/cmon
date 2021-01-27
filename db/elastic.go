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
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Bool("insecure_ssl")},
	}
	c = &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	host   = config.Str("host")
	port   = config.Str("port")
	https  = config.Bool("https")
	user   = config.Str("username")
	pass   = config.Str("password")
	scheme = "http"
)

const ct = "application/json;charset=UTF-8"

// Elastic makes POST HTTP call to ElasticSearch DB
func Elastic(d *data.DiscoverJSON) {
	if https {
		scheme = "https"
	}
	buf := new(bytes.Buffer)
	_, err := buf.WriteString(scheme)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString("://")
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString(host)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString(":")
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString(port)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString("/servers/_doc/")
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	_, err = buf.WriteString(config.Str("hostid"))
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	url := buf.String()

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(d)
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
