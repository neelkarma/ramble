package main

import (
	"os"

	"github.com/neelkarma/ramble/checks"
	"github.com/neelkarma/ramble/cmd"
	"github.com/neelkarma/ramble/constants"
)

func main() {
	cmd.Execute()
	if checks.TempFileExists() {
		os.Remove(constants.TempFilePath)
	}
}
