// Package boop contains library functions for boop.
package boop

import "os"

// Boop is the main functionality to create a file or directory.
type Boop interface {
	// Executable sets whether the created file should be executable. The default is false.
	Executable(executable bool)

	// Boop creates the named file or directory.
	Boop(name string) error
}

// New returns a new Boop.
func New() Boop {
	return &boop{}
}

// boop is the concrete implementation of Boop.
type boop struct {
	// executable tracks whether or not the file should be executable.
	executable bool
}

// Boop implements Boop.
func (boop boop) Boop(name string) error {
	var err error
	if name == "" {
		return ErrNameEmpty
	}
	err = createContainerDir(name)
	if err != nil {
		return err
	}
	if isDir(name) {
		return createDir(name)
	}
	perm := defaultPerm
	if boop.executable {
		perm |= executePerm
	}
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	// NOTE If the file already existed, OpenFile would not set the execute permission.
	// TODO Avoid redundancy by catching an already exists error and using that as a hint
	//		to set permissions.
	err = addExecutablePerm(f)

	return err
}
