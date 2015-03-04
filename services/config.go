package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Development ConfigSection
}

type ConfigSection struct {
	DbConnectionString string
	APIServer          string
}

var globalConfig Config

func init() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error", err)
	}
	json.Unmarshal(file, &globalConfig)
}

func GetConfig() *ConfigSection {
	return &globalConfig.Development
}
