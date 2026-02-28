//go:build unix

package boop

import "testing"

func TestExecutable(t *testing.T) {
	b := New()
	b.Executable(true)
	bb := b.(*boop)
	if !bb.executable {
		t.Fatalf(`Executable(true) should set executable to true`)
	}
}
