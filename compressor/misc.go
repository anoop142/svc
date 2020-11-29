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

// PrintLogo prints ascii
func PrintLogo() {
	const logo = `
	 ___              _      __   ___    _
	/ __(_)_ __  _ __| |___  \ \ / (_)__| |___ ___
	\__ \ | '  \| '_ \ / -_)  \ V /| / _‘ / -_) _ \
	|___/_|_|_|_| .__/_\___|   \_/ |_\__,_\___\___/
	  ___
	/ __|___ _ __  _ __ _ _ ___ _________ _ _
	|(__/ _ \ '  \| '_ \ '_/ -_|_-<_-< _ \ '_|
	\___\___/_|_|_| .__/_| \___/__/__|___/_|                  
	     Anoop S, EEE MACE,  ver 3.6
 `

	fmt.Println(logo)
}
