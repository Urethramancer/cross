package cross

import (
	"os/user"
)

func IsRoot() bool {
	u, err := user.Current()
	if err != nil {
		return false
	}

	if u.Name == "root" || u.Name == "Administrator" || u.Name == "System Administrator" {
		return true
	}

	return false
}
