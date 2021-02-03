package compressor

import (
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
		color.Question.Printf("Do you want to resize to 720p?\n")
		color.Bold.Printf("[Y/n] :")
		fmt.Scanf("%s\n", &downscaleOption)
	}

	if downscaleOption == "N" || downscaleOption == "n" {
		return false
	}

	// if downscaleOption is empty make it Y
	downscaleOption = "Y"
	return true

}

// checkRes checks if the video is higher than specified res dWidth x dHeight.
func checkRes(InputFile string) bool {

	var mediainfo string
	var width int64
	var height int64
	const (
		dWidth  = 1280
		dHeight = 720
	)

	if runtime.GOOS == "windows" {
		mediainfo = mediainfoWin
	} else {
		mediainfo = "mediainfo"
	}

	out, err := exec.Command(mediainfo, `--Inform=Video;%Width%x%Height%`, "Input/"+InputFile).Output()
	check(err)
	fmt.Sscanf(string(out), "%vx%v", &width, &height)

	if width > dWidth && height > dHeight {
		return resDialog()
	}
	return false

}

// IsScreenRecord asks the user if the source video is screen record
func IsScreenRecord() bool {
	opt := "Y"
	fmt.Println()

	color.Question.Printf("Is source a screen recording?\n")
	color.Bold.Printf("[Y/n] :")

	fmt.Scanf("%s\n", &opt)
	if opt == "N" || opt == "n" {
		return false
	}
	return true
}
