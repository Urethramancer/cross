package cross

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// SetBasePath builds the appropriate path for saving program-specific directories.
func SetBasePath() {
	if IsRoot() {
		basepath = "/etc"
		return
	}

	u, err := user.Current()
	if err != nil {
		return
	}

	dir := filepath.Join(u.HomeDir, "Library", "Application Support")
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return
		}
	}

	basepath = dir
}

// SetConfigPath builds the path to the command line program config directory.
func SetConfigPath(program string) {
	program = strings.Replace(program, " ", "", -1)
	dir := ""

	if IsRoot() {
		dir = filepath.Join(basepath, program)
	} else {
		dir = filepath.Join(basepath, program)
	}

	if !Exists(dir) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			configpath = basepath
			return
		}
	}

	configpath = dir
}

// SetServerConfigPath builds the path to the server config directory, bypassing automatically selected paths.
func SetServerConfigPath(program string) {
	SetConfigPath(program)
}

// GetServerConfigName gets the correct full path of the configuration file for servers.
func ServerConfigName(filename string) string {
	return ConfigName(filename)
}
