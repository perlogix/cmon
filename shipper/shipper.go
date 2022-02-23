package shipper

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
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Bool("api_insecure_ssl")},
	}
	c = &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}
	host  = config.Str("api_host")
	port  = config.Str("api_port")
	https = config.Bool("api_https")
	user  = config.Str("api_username")
	token = config.Str("api_token")
)

const ct = "application/json;charset=UTF-8"

// Ship makes POST HTTP call to Paradrop API
func Ship(d *data.DiscoverJSON) {
	var scheme string
	if https {
		scheme = "https"
	} else {
		scheme = "http"
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

	if user != "" && token != "" {
		r, err := http.NewRequest("POST", url, body)
		if err != nil {
			log.Printf("Error: %s\n", err)
			return
		}
		r.Header.Add("Content-Type", ct)
		r.SetBasicAuth(user, token)
		_, err = c.Do(r)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}
}
