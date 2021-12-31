package db

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/perlogix/cmon/config"
	"github.com/perlogix/cmon/data"
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
