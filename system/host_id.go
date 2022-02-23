package system

import (
	"fmt"

	"github.com/denisbrodbeck/machineid"
)

// GetHostID gets machine ID or UUID V4
func GetHostID() string {
	id, err := machineid.ProtectedID("cmon")
	if err != nil {
		fmt.Println(err)
		// Put UUID v4 code here
	}
	return id
}
