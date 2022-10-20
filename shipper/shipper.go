package shipper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
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
	url   = config.Str("api_url")
	user  = config.Str("api_username")
	token = config.Str("api_token")
)

const ct = "application/json"

// Ship makes POST HTTP call to Paradrop API
func Ship(d *data.DiscoverJSON) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(d)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	r.Header.Add("Content-Type", ct)
	r.Header.Add("X-Paradrop-Token", token)
	r.Header.Add("X-Paradrop-Email", user)

	resp, err := c.Do(r)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(respBody))
}
