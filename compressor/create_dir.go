package compressor

import "os"

// CreateDir creates Input ad Output folder if not exist
func CreateDir() {
	var err error

	err = os.MkdirAll("Input", 0777)
	check(err)

	err = os.MkdirAll("Ouput", 0777)
	check(err)
}
