package main

import (
	"encoding/json"
	"io"
	"os"
)

type config struct {
	UploadFilename string `json:"uploadFilename"`
	ServerAddress  string `json:"serverAddress"`
	Timezone       string `json:"timezone"`
}

func (conf *config) LoadConfig() error {
	data, err := os.Open("./settings.json")

	if err != nil {
		return err
	}

	byteValue, _ := io.ReadAll(data)

	json.Unmarshal(byteValue, conf)

	if conf.UploadFilename == "" {
		conf.UploadFilename = "noname.mp4"
	}

	if conf.ServerAddress == "" {
		conf.ServerAddress = "https://iotdevserver.inhandplus.com"
	}

	if conf.Timezone == "" {
		conf.Timezone = "Asia/Seoul"
	}
	return nil
}

func (conf *config) SaveConfig() error {
	data, err := json.Marshal(conf)

	if err != nil {
		return err
	}

	err = os.WriteFile("./settings.json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}