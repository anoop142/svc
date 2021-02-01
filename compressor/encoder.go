package compressor

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gookit/color"
)

// Compress - main encoding function
func Compress(filesToEncode []string, crf string) {

	var clearCmd string
	var ffmpeg string
	var tune = "stillimage"

	if runtime.GOOS == "windows" {
		clearCmd = "cls"
		ffmpeg = ffmpegWin
	} else {
		ffmpeg = "ffmpeg"
		clearCmd = "clear"
	}
	screenRecord := IsScreenRecord()

	for _, inputFile := range filesToEncode {
		outputFile := inputFile

		// Default to mkv if not either of these extensions
		if (strings.HasSuffix(inputFile, ".mp4") || strings.HasSuffix(inputFile, ".mkv")) == false {
			fileName := strings.Split(inputFile, ".")
			ext := fileName[len(fileName)-1]
			outputFile = strings.TrimSuffix(inputFile, ext) + "mkv"
		}
		mainParams := []string{"-loglevel", "quiet", "-stats", "-y", "-i", "Input/" + inputFile, "-pix_fmt", "yuv420p"}

		if !screenRecord {
			// Shot on Camera, not a screen record
			tune = "film"

			if checkRes(inputFile) {
				color.Warn.Println("Downscaling to  720p")
				mainParams = append(mainParams, "-vf", "scale=1280:720")
			}
		}

		videoParams := []string{"-c:v", "libx264", "-preset", x264Preset, "-tune", tune, "-crf", crf, "-r", fps, "-x264-params", x264Params}
		audioParams := []string{"-c:a", "libopus", "-b:a", opusBitrate, "-vbr", "on", "-compression_level", "10", "-frame_duration", "60", "-application", "audio", "-strict", "-2", "Ouput/" + outputFile}

		encodeCmd := append(mainParams, videoParams...)
		encodeCmd = append(encodeCmd, audioParams...)

		cmd := exec.Command(ffmpeg, encodeCmd...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		color.Info.Println("Compressing ", inputFile)

		err := cmd.Run()

		check(err)

		cmd = exec.Command(clearCmd)
		cmd.Stdout = os.Stdout
		cmd.Run()
		color.Info.Println(inputFile, "compressed successfully!")
	}
	color.Success.Printf("\nFinished!")
}
