package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eventials/go-tus"
)

func uploading(config *config) error {
	metadata := map[string]string{
		"filename": strings.Split(config.UploadFilename, "/")[0],
		"filetype": "mp4",
		"timezone": config.Timezone,
	}

	f, err := os.Open(config.UploadFilename)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	defer f.Close()

	// create the tus client.
	client, err := tus.NewClient(config.ServerAddress+"/files/", nil)

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
	// start the uploading process.
	err = uploader.Upload()

	if err != nil {
		sbox.AddLine(err.Error())
	}

	sbox.AddLine("Uplaod Completed")
	return nil
}
