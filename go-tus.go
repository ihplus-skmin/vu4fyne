package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eventials/go-tus"
)

func (v *uploadData) go_tus_upload(w *Widgets, garage tus.Store) (err error) {
	fi, err := v.fp.Stat()

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	config := tus.DefaultConfig()

	chunk, err := strconv.ParseInt(w.ChunkSizeEntry.Text, 10, 64)

	if err != nil {
		sbox.AddLine("Chunk size error")
		return
	}

	config.ChunkSize = chunk * 1024 * 1024
	config.Resume = true
	config.Store = garage

	client, err := tus.NewClient(v.url+"/files/", config)

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	fingerprint := fmt.Sprintf("%s-%d-%s", fi.Name(), fi.Size(), fi.ModTime())

	upload := tus.NewUpload(v.fp, fi.Size(), v.metadata, fingerprint)

	uploader, err := client.CreateUpload(upload)

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	sbox.AddLine("Upload started")

	w.Progress.Max = float64(fi.Size())

	if client.Config.ChunkSize < fi.Size() {
		w.Progress.Hidden = false
	} else {
		w.Progress.Hidden = true
	}

	for uploader.Offset() < fi.Size() {
		err = uploader.UploadChunck()

		if err != nil {
			sbox.AddLine(err.Error())
			return
		}

		w.Progress.SetValue(float64(uploader.Offset()))

		if client.Config.ChunkSize < fi.Size() && w.TargetServer.Text != "Dev Server" {
			time.Sleep(3 * time.Second)
		}
	}

	garage.Delete(fingerprint)

	w.Progress.SetValue(float64(fi.Size()))

	sbox.AddLine("Upload complete")
	return nil
}

func (v *uploadData) ResumeUpload(garage tus.Store) (err error) {
	fi, err := v.fp.Stat()

	if err != nil {
		sbox.AddLine(err.Error())
		return
	}

	config := tus.DefaultConfig()

	config.Resume = true
	config.Store = garage

	client, err := tus.NewClient(v.url+"/files/", config)

	if err != nil {
		return
	}

	upload := tus.NewUpload(v.fp, fi.Size(), v.metadata, "fingerprint")

	uploader, err := client.ResumeUpload(upload)

	if err != nil {
		return
	}

	for uploader.Offset() < fi.Size() {
		err = uploader.UploadChunck()

		if err != nil {
			return
		}
	}

	return nil
}
