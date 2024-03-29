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
	ChunkSize      string `json:"chunkSize"`
}

func (conf *config) LoadConfig() error {
	data, err := os.Open("./settings.json")

	if err != nil {
		_, err := os.Create("./settings.json")

		if err != nil {
			sbox.AddLine(err.Error())
			return err
		}

		data, err := json.Marshal(conf)

		if err != nil {
			sbox.AddLine(err.Error())
			return err
		}

		err = os.WriteFile("./settings.json", data, 0644)

		if err != nil {
			sbox.AddLine(err.Error())
			return err
		}
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

	if conf.ChunkSize == "" {
		conf.ChunkSize = "40"
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
