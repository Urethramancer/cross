package cross

import "path/filepath"

var basepath string
var configpath string

func init() {
	SetBasePath()
}

func BasePath() string {
	return basepath
}

func ConfigPath() string {
	return configpath
}

// GetConfigName gets the correct full path of the configuration file for command line tools.
func ConfigName(filename string) string {
	return filepath.Join(configpath, filename)
}
