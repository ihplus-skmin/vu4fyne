package main

import (
	"os"
	"strings"
)

func uploading(config *config) error {
	filename := strings.Split(config.UploadFilename, "/")

	metadata := map[string]string{
		"filename": filename[len(filename)-1],
		"filetype": "mp4",
		"timezone": config.Timezone,
	}

	f, err := os.Open(config.UploadFilename)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	//defer f.Close()

	var url string

	switch config.ServerAddress {
	case "Dev Server":
		url = "https://iotdevserver.inhandplus.com"
	case "Test Server":
		url = "https://iottestserver.inhandplus.com"
	case "Release Server":
		url = "https://iotserver.inhandplus.com"
	default:
		url = config.ServerAddress
	}

	fi, err := f.Stat()

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	err = tusOptions(url)

	if err != nil {
		sbox.AddLine("Server failure")
		return err
	}

	sbox.AddLine("Server exam. passed")

	location, err := tusPost(url, fi.Size(), metadata)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	sbox.AddLine("Location suuceefully aquired. \nUpload started")

	err = tusPatch(url, location, f, fi.Size(), config.UploadFilename)

	if err != nil {
		sbox.AddLine("Upload failure")
		return err
	}

	sbox.AddLine("Upload succeeded")

	return nil
}
