package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ProtocolVersion = "1.0.0"
)

func tusPost(address string, fileSize int64, metadata map[string]string) (location string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", address+"/files/", nil)

	if err != nil {
		sbox.AddLine("Making new POST request failed")
		return "", err
	}

	req.Header.Add("Tus-Resumable", ProtocolVersion)
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Content-Offset", "0")
	req.Header.Add("Upload-Length", strconv.FormatInt(fileSize, 10))
	req.Header.Add("Upload-Metadata", EncodedMetadata(metadata))

	res, err := client.Do(req)

	if err != nil {
		sbox.AddLine("POST request failed")
		return "", err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		location = res.Header.Get("Location")
	case 412:
		return "", errors.New("protocol version mismatch")
	case 413:
		return "", errors.New("upload body is to large")
	default:
		return "", fmt.Errorf("client error: status code %d", res.StatusCode)
	}

	return location, nil
}

func tusOptions(url string) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("OPTIONS", url+"/files/", nil)

	if err != nil {
		sbox.AddLine("Making new request failed")
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		sbox.AddLine("Options request failed")
		return
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200, 204:
		return nil
	case 404:
		errMsg := "server doesn't support tus protocol"
		sbox.AddLine(errMsg)
		return errors.New(errMsg)
	default:
		errMsg := fmt.Sprintf("client error: status code %d", res.StatusCode)
		sbox.AddLine(errMsg)
		return errors.New(errMsg)
	}
}

func tusPatch(url string, location string, fp *os.File, fileSize int64, filename string) (err error) {
	sarray := strings.Split(location, "/")

	hash := sarray[len(sarray)-1]

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if err != nil {
		sbox.AddLine(err.Error())
		return err
	}

	part, err := writer.CreateFormFile("file", filepath.Base(filename))

	if err != nil {
		return
	}

	_, err = io.Copy(part, fp)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("PATCH", url+"/files/"+hash, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/offset+octet-stream")
	req.Header.Add("Tus-Resumable", "1.0.0")
	req.Header.Add("Content-Length", strconv.FormatInt(int64(fileSize), 10))
	req.Header.Add("Upload-Offset", "0")

	req.Header.Set("X-HTTP-Method-Override", "PATCH")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 204:
		if newOffset, err := strconv.ParseInt(res.Header.Get("Upload-Offset"), 10, 64); err == nil {
			fmt.Println(newOffset)
			return nil
		} else {
			return err
		}
	case 409:
		return errors.New("offset mismatch")
	case 412:
		return errors.New("version mismatch")
	case 413:
		sbox.AddLine("request entity too large")
		return errors.New("large upload")
	default:
		return errors.New("client error")
	}
}
