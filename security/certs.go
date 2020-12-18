package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/yeticloud/yeti-discover/data"
)

// ExpiredCerts detects if certs are expiring within 30 days
func ExpiredCerts(d *data.DiscoverJSON) {
	if runtime.GOOS != "windows" {
		if runtime.GOOS != "darwin" {
			cmd := `find /etc/ssl /etc/pki /usr/local/share/certs /etc/openssl /var/ssl /usr/local/share/ca-certificates /etc/apache2 /etc/httpd /etc/nginx -type f \( -iname \*.pem -o -iname \*.crt -o -iname \*.cert \)  -printf "%p " -exec openssl x509 -noout -checkend 2592000 -in {} 2>/dev/null \; 2>/dev/null | grep -v "not expire"`
			out, err := exec.Command("/bin/sh", "-c", cmd).Output()
			if err != nil {
				return
			}

			var certs = strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")

			d.ExpiredCerts = certs
		}
	}
}
