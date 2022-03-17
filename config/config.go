package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	defaultConfigPath = "config/"
	defaultConfigName = "app.json"
	envKeyProfile     = "CONFIG_PROFILE"
)

func Load(configStruct interface{}) error {
	return LoadFromWithProfile(defaultConfigPath, CurrentProfile(), configStruct)
}

func LoadFromWithProfile(configDirectory, profile string, configStruct interface{}) error {
	configFileName := determineFileName(profile)
	return loadFromFile(configDirectory, configFileName, configStruct)
}

func CurrentProfile() string {
	return os.Getenv(envKeyProfile)
}

func MustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("No ENV value defined for key '%s'", key))
	}
	return value
}

func loadFromFile(configDirectory, configFileName string, configStruct interface{}) error {
	configFilePath := path.Join(configDirectory, configFileName)
	if !fileExists(configFilePath) {
		return fmt.Errorf("no config file found at '%s'", configFilePath)
	}

	return readJSONConfigFile(configFilePath, configStruct)
}

func readJSONConfigFile(fileName string, configStruct interface{}) error {
	cfgFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer cfgFile.Close()

	data, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, configStruct)
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func determineFileName(profile string) string {
	if profile == "" {
		return defaultConfigName
	}
	return "app-" + profile + ".json"
}
