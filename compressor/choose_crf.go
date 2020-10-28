package compressor

import (
	"fmt"

	"github.com/gookit/color"
)

//ChooseCRF displays a dialog to choose CRF value
func ChooseCRF() string {
	crf := "32.0"
	fmt.Println()

	color.Note.Printf("Choose a CRF value between 28 and 35 or empty for default value(%v)\n", crf)
	color.Question.Printf("CRF value :")

	fmt.Scanf("%s\n", &crf)
	return crf
}
