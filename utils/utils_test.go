package utils_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/chayut-eam/test-golang-hugeman/utils"
)

func TestGetRelativePath(t *testing.T) {
	fmt.Println(utils.GetRelativePath())
}

func TestConvertImageToBase64(t *testing.T) {
	//success
	currentPath, _ := utils.GetRelativePath()
	imgPath := filepath.Join(currentPath, "utils", "test.jpg")
	image, _ := os.Open(imgPath)
	defer image.Close()
	fmt.Println(utils.ConvertImgToBase64(image))

	//error
	currentPath, _ = utils.GetRelativePath()
	imgPath = filepath.Join(currentPath, "xxxxx", "test.jpg")
	image, _ = os.Open(imgPath)
	defer image.Close()
	fmt.Println(utils.ConvertImgToBase64(image))
}

func TestGetNow(t *testing.T) {
	fmt.Println(utils.Now())
}
