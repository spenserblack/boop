package boop

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateContainerDir(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf(`Could not create a temp directory: %v`, err)
	}
	defer os.RemoveAll(tempdir)

	const target string = "deeply/nested/file.go"
	fullTarget := filepath.Join(tempdir, target)
	fullTargetDir := filepath.Dir(fullTarget)
	err = createContainerDir(fullTarget)
	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}
	info, err := os.Lstat(fullTargetDir)
	if err != nil {
		t.Fatalf(`Could not get file info: %v`, err)
	}
	if !info.IsDir() {
		t.Errorf(`fullTargetDir should be a directory`)
	}
}
