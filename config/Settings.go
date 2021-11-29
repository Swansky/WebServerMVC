package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Settings struct {
	Database struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"databaseName"`
		Port         int    `yaml:"port"`
	} `yaml:"database"`
}

var settings *Settings

func LoadSettings() *Settings {
	f, err := os.Open("settings.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Settings
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	settings = &cfg
	return settings
}

func GetSettings() *Settings {
	return settings

}
