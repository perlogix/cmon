package security

import (
	"runtime"
	"strings"

	"github.com/perlogix/cmon/data"
	"github.com/perlogix/cmon/util"
)

// ExpiredCerts detects if certs are expiring within 30 days
func ExpiredCerts(d *data.DiscoverJSON) {
	expireCerts := []string{}

	if runtime.GOOS == "linux" {

		certOut, err := util.Cmd(`find /etc/ssl /etc/pki /usr/local/share/certs /etc/openssl /var/ssl /usr/local/share/ca-certificates /etc/apache2 /etc/httpd /etc/nginx -type f \( -iname \*.pem -o -iname \*.crt -o -iname \*.cert \)  -printf "%p " -exec openssl x509 -noout -checkend 2592000 -in {} 2>/dev/null \; 2>/dev/null | grep -v "not expire"`)
		if err != nil {
			d.ExpiredCerts = expireCerts
			return
		}

		expireCerts = append(expireCerts, strings.Split(strings.TrimSuffix(string(certOut), "\n"), "\n")...)
	}

	d.ExpiredCerts = expireCerts
}
