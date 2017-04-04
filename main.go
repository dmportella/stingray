package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	_ "time"
)

func main() {
	fmt.Println("Hello stingray")

	config := Config{}

	if currentFolder, ok1 := getCurrentFolder(); ok1 == nil {
		commandSettingsPath := path.Join(currentFolder, "config.json")

		/*if data, ok2 := toJSON(&config); ok2 == nil {
			if ok3 := saveToFile(data, commandSettingsPath); ok3 == nil {
			}
		}*/

		if data, ok2 := readFromFile(commandSettingsPath); ok2 == nil && data != nil {
			fromJSON(data, &config)
		}
	}

	if client, err := NewClient(config); err == nil {
		if actions, err := client.GetActionList("/api/tm"); err == nil {
			fmt.Println(actions)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

// DefaultPermissions Permissions for file and directory creations.
const (
	defaultPermissions = 0666
)

// GetCurrentFolder Returns the
func getCurrentFolder() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}

// FileOrDirectoryExists Check is a file or directory exists.
func fileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// SaveToFile Saves a byte array to a path.
func saveToFile(data []byte, path string) error {
	return ioutil.WriteFile(path, data, defaultPermissions)
}

// ReadFromFile Safe reads file content.
func readFromFile(filepath string) ([]byte, error) {
	if ok, err := fileOrDirectoryExists(filepath); err != nil || !ok {
		return nil, err
	}

	return ioutil.ReadFile(filepath)
}

// FromJSON converts json object representation from a byte array into golang struct.
func fromJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return errors.New("json decoding error")
	}

	return nil
}

// ToJSON coverts a object into it's json representation and returns a byte array.
func toJSON(v interface{}) (data []byte, err error) {
	data, err = json.Marshal(v)
	if err != nil {
		return
	}
	return
}
