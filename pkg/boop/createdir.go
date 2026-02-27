package boop

import "os"


// createDir creates a directory.
func createDir(dirname string) error {
	if dirname == "" || dirname == "." {
		return nil
	}

	return os.MkdirAll(dirname, directoryPerm)
}
