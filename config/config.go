package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Application struct {
		Name                string `yaml:"name"`
		Logdir              string `yaml:"logDir"`
		ODApplicationDBUrl  string `yaml:"odApplicationDBUrl"`
		ODApplicationDBName string `yaml:"odApplicationDBName"`
	} `yaml:"application"`
}

var appConfig *Config

func GetAppConfig() *Config {
	if appConfig == nil {
		appConfig = loadConfiguration("conf/config.yaml")
	}
	return appConfig
}

func loadConfiguration(fileName string) *Config {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Cannot read app configuraion : ", fileName, err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalln("Error Cannot initial app configuraion : ", fileName, err)
	}
	return &cfg
}
