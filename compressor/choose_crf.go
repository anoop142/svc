package compressor

import (
	"fmt"
	"strconv"

	"github.com/gookit/color"
)

const defaultCRF = "32"

//ChooseCRF displays a dialog to choose CRF value
func ChooseCRF() string {
	var opt int

	fmt.Println()
	color.Note.Println("Choose a compression value between 28 and 35 or empty for default value(%s)\n", defaultCRF)
	color.Question.Printf("Compression value(CRF) :")

	fmt.Scanf("%d", &opt)

	if opt == 0 {
		return defaultCRF
	}
	return strconv.Itoa(opt)

}
