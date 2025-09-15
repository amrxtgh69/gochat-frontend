package handlers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// get rc file path (~/.gochatrc)
func getRCFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gochatrc")
}

// read rc file into map
func readRCFile() (map[string]string, error) {
	data := make(map[string]string)

	filePath := getRCFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil // return empty map if not exists
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			data[parts[0]] = parts[1]
		}
	}
	return data, nil
}

// write rc map back to file
func writeRCFile(data map[string]string) error {
	filePath := getRCFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range data {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return err
		}
	}
	return nil
}

// set or update a key
func setRCValue(key, value string) error {
	data, err := readRCFile()
	if err != nil {
		return err
	}
	data[key] = value
	return writeRCFile(data)
}

// get a key
func GetRCValue(key string) (string, error) {
	data, err := readRCFile()
	if err != nil {
		return "", err
	}
	return data[key], nil
}

// delete a key
func deleteRCValue(key string) error {
	data, err := readRCFile()
	if err != nil {
		return err
	}
	if _, exists := data[key]; exists {
		delete(data, key)
		return writeRCFile(data)
	}
	return nil // key doesn't exist, nothing to do
}
