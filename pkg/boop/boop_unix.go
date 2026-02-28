//go:build unix

package boop

// Executable implements Boop.
func (boop *boop) Executable(executable bool) {
	boop.executable = executable
}
