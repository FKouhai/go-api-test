package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct { //config struct which contains the bot Token and its Prefix
	MongoHost       string `json:"mongoHost"`
	Port            string `json:"Port"`
	MongoDB         string `json:"mongoDB"`
	MongoCollection string `json:"mongoCollection"`
}

func ReadConfig() *Config { //Function that reads the config file using json unmarshal
	var config *Config
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return config
}
