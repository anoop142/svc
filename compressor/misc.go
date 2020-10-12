package compressor

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// WaitFun a hacky function to prevent closing of window, mainly for Windows
func WaitFun() {
	fmt.Println()
	fmt.Print("Press Enter key to exit..")
	fmt.Scanln()
	os.Exit(0)
}

func PrintLogo() {
	const logo = `
======================================
	  Video Compressor  
	    ver 2.5
	Anoop S 
======================================
`
	fmt.Println(logo)
}
