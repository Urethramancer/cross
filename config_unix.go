// +build linux dragonfly freebsd netbsd openbsd solaris

package cross

import (
	"os"
	"os/user"
	"path/filepath"
)

// SetBasePath builds the appropriate path for saving program-specific directories.
func SetBasePath() {
	u, err := user.Current()
	if err != nil {
		return
	}

	basepath = u.HomeDir
}

// SetConfigPath builds the path to the command line program config directory.
func SetConfigPath(program string) {
	dir := filepath.Join(basepath, "."+program)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			configpath = basepath
			return
		}
	}

	configpath = dir
}

// GetServerConfigName gets the correct full path of the configuration file for servers.
func ServerConfigName(program, filename string) string {
	dir := filepath.Join("/etc", program)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return filename
		}
	}

	return filepath.Join(dir, filename)
}
