package config

import (
	"fmt"
	"sync"
	"os"
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

func CreateConfig() *Config {
	exists, err := checkForConfig()
	if err != nil {
		fmt.Println("error checking for config:", err)
		return nil
	}
	if exists {
		fmt.Println("config already exists")
		return nil
	}
	config := &Config{}
	_, err = os.Create("config.ini")
	if err != nil {
		fmt.Println("error creating config:", err)
		return nil
	}
	ini.Set("fileId", "")


	
	return config
}
func checkForConfig() (bool, error) {
	_, err := os.Stat("config.ini")
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}

}