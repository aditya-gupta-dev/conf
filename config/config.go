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

func WriteConfig[T any](configModel T, appName, filename string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	configFilePath := filepath.Join(cacheDir, appName, filename)

	if isFileExists(configFilePath) {
		jsonData, err := json.MarshalIndent(configModel, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(configFilePath, jsonData, 0644); err != nil {
			return err
		}
	} else {
		return os.ErrNotExist
	}

	return nil
}

func ModifyConfig[T any](configModel *T, appName, filename string) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	configFilePath := filepath.Join(cacheDir, appName, filename)

	if !isFileExists(configFilePath) {
		return os.ErrNotExist
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, configModel); err != nil {
		return err
	}

	return nil
}

func GetConfig[T any](appName, filename string) (T, error) {
	var configModel T

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return configModel, err
	}

	configFilePath := filepath.Join(cacheDir, appName, filename)

	if !isFileExists(configFilePath) {
		return configModel, os.ErrNotExist
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return configModel, err
	}

	if err := json.Unmarshal(data, &configModel); err != nil {
		return configModel, err
	}

	return configModel, nil
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
