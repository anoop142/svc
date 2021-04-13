package compressor

import "os"

// CreateDir creates Input ad Output folder if not exist
func CreateDir() {
	var err error

	err = os.MkdirAll(InputDir, 0777)
	check(err)

	err = os.MkdirAll(OutputDir, 0777)
	check(err)
}
