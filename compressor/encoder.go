package compressor

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/gookit/color"
)

const opusBitrate = "64k"

// Insert downscaling params
func insertScaleParams(encodeCmd []string) []string {

	scaleParams := []string{"-vf", "scale=1280:720"}

	return append(encodeCmd[:6], append(scaleParams, encodeCmd[6:]...)...)
}

// Compress - main encoding function
func Compress(filesToEncode []string, crf string) {

	ffmpeg := "ffmpeg"
	clearCmd := "clear"

	if runtime.GOOS == "windows" {
		clearCmd = "cls"
		ffmpeg = "bin/ffmpeg.exe"
	}
	for _, f := range filesToEncode {

		mainParams := []string{"-loglevel", "quiet", "-stats", "-y", "-i", "Input/" + f}

		videoParams := []string{"-c:v", "libx264", "-preset", "slow", "-crf", crf, "-r", "25", "-x264-params", "ref=6:qpmin=10:qpmax=51:me=umh:bframes=6"}
		audioParams := []string{"-c:a", "libopus", "-b:a", opusBitrate, "-vbr", "on", "-compression_level", "10", "-frame_duration", "60", "-application", "audio", "-strict", "-2", "Ouput/" + f}

		encodeCmd := append(mainParams, videoParams...)
		encodeCmd = append(encodeCmd, audioParams...)

		if checkRes(f) {
			// insert downscaling param
			color.Warn.Println("Downscaling to  720p")
			insertScaleParams(encodeCmd)
		}

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
