package boop

import "io/fs"

const (
	// defaultPerm gives user, group, and others read/write permission.
	defaultPerm fs.FileMode = 0o666

	// executePerm gives user, group, and others execute permission.
	executePerm fs.FileMode = 0o111

	// directoryPerm gives the user full access, group and others read/execute permission.
	directoryPerm fs.FileMode = 0o755
)
