package boop

import "testing"

func TestNew(t *testing.T) {
	b := New().(*boop)
	if b.executable {
		t.Errorf(`The default value of executable should be false`)
	}
}
