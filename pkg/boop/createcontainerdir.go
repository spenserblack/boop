package boop

import "path/filepath"

// createContainerDir creates a directory for a file.
func createContainerDir(filename string) error {
	dir := filepath.Dir(filename)
	return createDir(dir)
}
