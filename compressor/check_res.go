package compressor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gookit/color"
)

var downscaleOption string

func resDialog() bool {

	// Initial selection
	if downscaleOption == "" {
		fmt.Println()
		color.Question.Println("Do you want to resize to 720p? ")
		color.Bold.Print("[Y/n] :")
		fmt.Scanf("%s", &downscaleOption)
	}

	if downscaleOption == "N" || downscaleOption == "n" {
		return false
	}

	// if downscaleOption is empty make it Y
	downscaleOption = "Y"
	return true

}

func checkRes(InputFile string) bool {

	var ffprobe string

	if runtime.GOOS == "windows" {
		ffprobe = "bin/ffprobe.exe"
	} else {
		ffprobe = "ffprobe"
	}

	var out bytes.Buffer

	type Mediastruct struct {
		Streams []struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"streams"`
	}

	var MediaInfo Mediastruct

	cmd := exec.Command(ffprobe, "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "json", "Input/"+InputFile)
	cmd.Stdout = &out

	err := cmd.Run()
	check(err)

	err = json.Unmarshal(out.Bytes(), &MediaInfo)
	check(err)
	for _, i := range MediaInfo.Streams {
		if i.Height > 720 && i.Width > 1280 {
			return resDialog()
		}
	}
	return false

}

// IsScreenRecord asks the user if the source video is screen record
func IsScreenRecord() bool {
	opt := "Y"
	fmt.Println()

	color.Question.Println("Is source a screen recording?")
	color.Bold.Print("[Y/n] :")

	fmt.Scanf("%s", &opt)
	if opt == "N" || opt == "n" {
		return false
	}
	return true
}
