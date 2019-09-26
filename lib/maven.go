package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ReplaceMavenTemplate takes filename and replaces the strings PACMAN_USER
// and PACMAN_PASS with username and password, respectively
func ReplaceMavenTemplate(filename, username, password string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	newContent := strings.ReplaceAll(string(content), "PACMAN_USER", username)
	newContent = strings.ReplaceAll(string(newContent), "PACMAN_PASS", password)
	return []byte(newContent), nil
}

// WriteMavenTemplate writes bytes to filename
func WriteMavenTemplate(bytes []byte, filename string) error {
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, bytes, 0600); err != nil {
		return err
	}
	fmt.Printf("File successfully written at: %s", filename)
	return nil
}

// GetDefaultMavenConfPath returns the absolute path of where
// maven's settings.xml is usually found by convention
func GetDefaultMavenConfPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".m2", "settings.xml"), nil
}
