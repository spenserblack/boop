package boop

// isDir returns true if the target should be a directory, false if it should be a file.
func isDir(target string) bool {
	lastRune := rune(target[len(target) - 1])
	for _, r := range slashes {
		if r == lastRune {
			return true
		}
	}
	return false
}
