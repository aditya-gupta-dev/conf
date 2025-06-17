package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func InitConfig[T any](configModel T, appName, filename string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	configFilePath := filepath.Join(cacheDir, appName, filename)
	if !isFileExists(configFilePath) {
		if err := os.MkdirAll(filepath.Join(cacheDir, appName), 0755); err != nil {
			return err
		}

		file, err := os.OpenFile(configFilePath, os.O_CREATE, 0644)
		if err != nil {
			return err
		}

		jsonData, err := json.MarshalIndent(configModel, "", "  ")
		if err != nil {
			return err
		}

		_, err = file.Write(jsonData)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil

}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
