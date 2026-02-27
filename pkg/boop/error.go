package boop

import "errors"

// ErrNotAFile is raised if the file to be created or modified is not a file, but
// actions specific to files should be taken.
var ErrNotAFile = errors.New("Action can only be taken on files")

// ErrNameEmpty is raised when an empty string is passed as the file name.
var ErrNameEmpty = errors.New("File name cannot be empty")
