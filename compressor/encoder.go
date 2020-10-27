package compressor

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/gookit/color"
)

const opusBitrate = "72k"

// Compress - main encoding function
func Compress(filesToEncode []string, crf string) {

	var clearCmd string
	var ffmpeg string

	if runtime.GOOS == "windows" {
		clearCmd = "cls"
		ffmpeg = "bin/ffmpeg.exe"
	} else {
		ffmpeg = "ffmpeg"
		clearCmd = "clear"
	}

	for _, f := range filesToEncode {

		mainParams := []string{"-loglevel", "quiet", "-stats", "-y", "-i", "Input/" + f}

		if checkRes(f) {
			// insert downscaling param
			color.Warn.Println("Downscaling to  1080p")
			mainParams = append(mainParams, "-vf", "scale=1920:1080")
		}

		videoParams := []string{"-c:v", "libx264", "-preset", "slow", "-crf", crf, "-r", "25", "-x264-params", "ref=6:qpmin=10:qpmax=51:me=umh:bframes=6"}
		audioParams := []string{"-c:a", "libopus", "-b:a", opusBitrate, "-vbr", "on", "-compression_level", "10", "-frame_duration", "60", "-application", "audio", "-strict", "-2", "Ouput/" + f}

		encodeCmd := append(mainParams, videoParams...)
		encodeCmd = append(encodeCmd, audioParams...)

		cmd := exec.Command(ffmpeg, encodeCmd...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		color.Info.Println("Compressing ", f)

		err := cmd.Run()

		check(err)

		cmd = exec.Command(clearCmd)
		cmd.Stdout = os.Stdout
		cmd.Run()
		color.Info.Println(f, "compressed successfully!")
	}
	color.Success.Printf("\nFinished!")
}
