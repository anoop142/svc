package compressor

import (
	"fmt"
	"io/ioutil"

	"github.com/gookit/color"
)

// ChooseInputFile list all files and takes int input
func ChooseInputFile() []string {
	var fileList []string
	var opt int
	files, err := ioutil.ReadDir("Input")
	check(err)

	if len(files) == 0 {
		color.Error.Println("No video files in Input folder")
		color.Note.Println("Put the video file in Input folder to compress! ")
		WaitFun()
	}

	i := 0

	for _, file := range files {

		fileList = append(fileList, file.Name())
		color.Bold.Printf("[%v]", i)
		fmt.Printf("%v\n", file.Name())
		i++
	}

	// If only one file
	if len(fileList) == 1 {
		return fileList
	}

	color.Bold.Printf("[%v]", i)
	fmt.Printf("All")

	fmt.Println(" ")
	color.Warn.Println("Choose a input file")
	fmt.Println("eg : 0")
	for {
		color.Question.Print("Enter selection number: ")
		fmt.Scanf("%d", &opt)
		if opt >= 0 && opt < len(fileList) {
			tmp := []string{fileList[opt]}
			return tmp
		} else if opt == len(fileList) {
			// encode all
			return fileList

		} else {
			color.Error.Println("Wrong selection")
			continue
		}
	}

}
