package boop

import "os"

// addExecutablePerm adds the executable permission to a file.
func addExecutablePerm(f *os.File) error {
	stat, err := f.Stat()
	if err != nil {
		return err
	}
	mode := stat.Mode()
	mode |= executePerm
	return f.Chmod(mode)
}
