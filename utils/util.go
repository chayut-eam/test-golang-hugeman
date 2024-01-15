package utils

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"time"
)

func Now() string {
	bangkokTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	currentTimeInBangkok := time.Now().In(bangkokTimeZone)
	now, _ := time.Parse(time.RFC3339, currentTimeInBangkok.Format(time.RFC3339))
	return now.Format(time.RFC3339)
}

func ConvertImgToBase64(imageFile io.Reader) (*string, error) {
	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return nil, err
	}

	// Encode the image bytes to base64
	base64Encoded := base64.StdEncoding.EncodeToString(imageData)
	return &base64Encoded, nil
}

func GetRelativePath() (string, error) {
	// Get the path to the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the current file
	dir := filepath.Dir(currentFile)
	dir = filepath.Dir(dir)

	return dir, nil
}
