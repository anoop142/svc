package compressor

import (
	"fmt"
	"io/ioutil"

	"github.com/gookit/color"
)

// ChooseInputFile list all files and takes int input
func ChooseInputFile() []string {
	var fileList []string

	files, err := ioutil.ReadDir("Input")
	check(err)

	fileCount := len(files)

	// default to batch mode
	opt := fileCount

	if fileCount == 0 {
		color.Error.Println("No video files in Input folder")
		color.Note.Println("Put the video file in Input folder to compress! ")
		WaitFun()
	}

	for i, file := range files {

		fileList = append(fileList, file.Name())
		fmt.Printf("[%v] %v\n", color.Bold.Render(i), file.Name())
	}

	// If only one file
	if fileCount == 1 {
		return fileList
	}

	fmt.Printf("\n[%v] All\n", color.Bold.Render(fileCount))

	fmt.Println(" ")
	color.Warn.Println("Choose a input file")
	fmt.Println("eg : 0")
	for {
		color.Question.Print("Enter selection number: ")
		fmt.Scanf("%d", &opt)
		if opt >= 0 && opt < fileCount {
			tmp := []string{fileList[opt]}
			return tmp
		} else if opt == fileCount {
			// encode all
			return fileList

		} else {
			color.Error.Println("Wrong selection")
			continue
		}
	}

}
