package main

import "github.com/anoop142/svc/compressor"

func main() {

	compressor.PrintLogo()

	compressor.CreateDir()

	compressor.Compress(compressor.ChooseInputFile(), compressor.ChooseCRF())

	compressor.WaitFun()

}
