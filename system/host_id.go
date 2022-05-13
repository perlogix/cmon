package system

import (
	"github.com/denisbrodbeck/machineid"
)

// GetHostID gets machine ID or UUID V4
func GetHostID() string {
	id, err := machineid.ProtectedID("cmon")
	if err != nil {
		return ""
	}
	return id
}
