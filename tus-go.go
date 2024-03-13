package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eventials/go-tus"
)

func (v *uploadData) go_tus_upload() (err error) {
	fi, err := v.fp.Stat()

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}
	// create the tus client.
	client, err := tus.NewClient(v.url+"/files/", nil)

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	// create an upload from a file.
	fingerprint := fmt.Sprintf("%s-%d-%s", fi.Name(), fi.Size(), fi.ModTime())

	upload := tus.NewUpload(v.fp, fi.Size(), v.metadata, fingerprint)

	// create the uploader.
	uploader, err := client.CreateUpload(upload)

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	// start the uploading process.
	sbox.AddLine("Upload started")

	for uploader.Offset() < fi.Size() {
		sbox.AddLine(strconv.FormatInt(int64(uploader.Offset()), 10))
		err = uploader.UploadChunck()

		if err != nil {
			sbox.AddLine(err.Error())
			return
		}
		time.Sleep(3 * time.Second)
	}

	sbox.AddLine("Upload complete")
	return nil
}
