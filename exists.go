package cross

import "os"

// Exists checks for the existence of a file or directory.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks for the existence of a directory.
// It will return false if a file with that name exists.
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return true
	}

	return false
}

// FileExists checks for the existence of a file.
// It will return false if a directory with that name exists.
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}
