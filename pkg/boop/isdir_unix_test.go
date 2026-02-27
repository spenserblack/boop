//go:build unix
package boop

import "testing"

func TestIsDir(t *testing.T) {
	tests := []struct {
		name string
		target string
		want bool
	}{
		{
			name: "trailing /",
			target: "foo/",
			want: true,
		},
		{
			name: "no trailing slash",
			target: "foo",
			want: false,
		},
		{
			name: `trailing \`,
			target: `foo\`,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDir(tt.target); got != tt.want {
				t.Fatalf(`isDir(%q) = %v, want %v`, tt.target, got, tt.want)
			}
		})
	}
}
