package main

import (
	"fmt"
	"os"

	"github.com/eventials/go-tus"
)

func uploading() {
	metadata := map[string]string{
		"filename": "21IHPA00006A_231013_161517_manual.mp4",
		"filetype": "mp4",
		"timezone": "Asia/Seoul",
	}

	f, err := os.Open("./21IHPA00006A_231013_161517_manual.mp4")

	if err != nil {
		fmt.Println("Hey Paniced")
		return
	}

	defer f.Close()

	// create the tus client.
	client, err := tus.NewClient("http://3.36.35.75:8080/files/", nil)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fi, err := f.Stat()

	if err != nil {
		return
	}

	fingerprint := fmt.Sprintf("%s-%d-%s", fi.Name(), fi.Size(), fi.ModTime())

	upload := tus.NewUpload(f, fi.Size(), metadata, fingerprint)

	// create the uploader.
	uploader, err := client.CreateUpload(upload)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// start the uploading process.
	uploader.Upload()
}
