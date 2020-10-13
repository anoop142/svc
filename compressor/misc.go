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
	 _____            __      _   ___    __                                                 
	/ __(_)_ _  ___  / /__   | | / (_)__/ /__ ___    
   _\ \/ /  ' \/ _ \/ / -_)  | |/ / / _  / -_) _ \  
  /___/_/_/_/_/ .__/_/\__/   |___/_/\_,_/\__/\___/  
			 /_/                                     
	_____                                       
   / ___/__  __ _  ___  _______ ___ ______  ____
  / /__/ _ \/  ' \/ _ \/ __/ -_|_-<(_-< _ \/ __/
  \___/\___/_/_/_/ .__/_/  \__/___/___|___/_/   
				/_/                             
    	Anoop S,  ver 2.6
 `

	fmt.Println(logo)
}
