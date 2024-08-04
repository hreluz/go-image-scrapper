package fileutil

import (
	"encoding/json"
	"os"
)

// SaveToFile saves an interface to a file as JSON.
func SaveToFile(filename string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads JSON data from a file into an interface.
func LoadFromFile(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// FileExists checks if a file exists and is not a directory.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DeleteFile(filename string) {
	if FileExists(filename) {
		os.Remove(filename)
	}
}
