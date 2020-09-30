package compressor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gookit/color"
)

func resDialog() bool {
	opt := "Y"
	fmt.Println()
	color.Question.Println("Do you want to resize to 720p ")
	color.Bold.Print("Y / N :")
	fmt.Scanf("%v", &opt)
	if opt[0] == 'N' || opt[0] == 'n' {
		return false
	}
	return true

}

func checkRes(InputFile string) bool {
	var ffprobe = "ffprobe"

	var out bytes.Buffer

	type Mediastruct struct {
		Streams []struct {
			Width        int `json:"width"`
			Height       int `json:"height"`
			SideDataList []struct {
			} `json:"side_data_list"`
		} `json:"streams"`
	}

	var MediaInfo Mediastruct

	if runtime.GOOS == "windows" {
		ffprobe = "bin/ffprobe.exe"
	}
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
