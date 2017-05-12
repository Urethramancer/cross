package cross

import "os"

// Exists checks for the existence of a file or directory.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
