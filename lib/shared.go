package lib

import "os"

// FileExists returns true if file exists BUT there may be cases
// where file exists but no read permissions lead to error
// (in which case false is returned)
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
