package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func EncodedMetadata(metadata map[string]string) string {
	var encoded []string

	for k, v := range metadata {
		encoded = append(encoded, fmt.Sprintf("%s %s", k, b64encode(v)))
	}

	return strings.Join(encoded, ",")
}

func b64encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
