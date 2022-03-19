package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/neelkarma/ramble/constants"
)

type Config struct {
	Passhash string
	EncSalt  string
	EncIv    string
}

func ReadConfig() *Config {
	file, err := os.Open(constants.ConfigPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &config)

	return &config
}

func WriteConfig(config *Config) {
	file, err := os.Create(constants.ConfigPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}
	file.Write(bytes)
}
