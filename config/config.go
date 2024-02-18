package config

import (
	"fmt"
	"sync"

	"github.com/gookit/ini/v2"
)

var (
	once sync.Once
	config *Config
)

type Config struct {
	FileId      string
	SecretsPath string
	DataFile    string
}

func initConfig() {
	config = &Config{}
	err := ini.LoadExists("config.ini")
	if err != nil {
		fmt.Printf("error loading config: %v", err)
	}
	config.FileId = ini.String("fileId")
	config.SecretsPath = ini.String("secretsPath")
	config.DataFile = ini.String("dataFile")
	if config.FileId == "" || config.SecretsPath == "" || config.DataFile == "" {
		 fmt.Printf("missing required config")
	}
}

func GetConfig() *Config {
	once.Do(initConfig)
	return config
}