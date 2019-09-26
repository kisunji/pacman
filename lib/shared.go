package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

// WriteToFile creates parent directories if they do not exist
func WriteToFile(bytes []byte, filename string) error {
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, bytes, 0600); err != nil {
		return err
	}
	fmt.Printf("File successfully written to: %s", filename)
	return nil
}
