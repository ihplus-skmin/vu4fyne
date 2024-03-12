package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eventials/go-tus"
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

	defer f.Close()

	var address string

	switch config.ServerAddress {
	case "Dev Server":
		address = "https://iotdevserver.inhandplus.com"
	case "Test Server":
		address = "https://iottestserver.inhandplus.com"
	case "Release Server":
		address = "https://iotserver.inhandplus.com"
	default:
		address = config.ServerAddress
	}
	// create the tus client.
	client, err := tus.NewClient(address+"/files/", nil)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	fi, err := f.Stat()

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	fingerprint := fmt.Sprintf("%s-%d-%s", fi.Name(), fi.Size(), fi.ModTime())

	upload := tus.NewUpload(f, fi.Size(), metadata, fingerprint)

	// create the uploader.
	uploader, err := client.CreateUpload(upload)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	sbox.AddLine("Upload started")
	// start the uploading process.
	err = uploader.Upload()

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	sbox.AddLine("Uplaod Completed")
	return nil
}
