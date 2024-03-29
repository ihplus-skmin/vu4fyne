package main

import (
	"os"
	"strings"

	"github.com/eventials/go-tus"
)

type uploadData struct {
	url            string
	uploadFilename string
	fp             *os.File
	metadata       map[string]string
}

func uploading(config *config, w *Widgets, g tus.Store) error {
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

	v := uploadData{
		url:            url,
		uploadFilename: config.UploadFilename,
		fp:             f,
		metadata:       metadata,
	}

	err = v.go_tus_upload(w, g)

	if err != nil {
		return err
	}

	return nil
}
