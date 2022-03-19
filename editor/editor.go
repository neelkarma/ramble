package editor

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/neelkarma/ramble/constants"
)

func Open(initialValue string) (string, error) {
	f, err := os.Create(constants.TempFilePath)

	f.WriteString(initialValue)

	if err != nil {
		return "", err
	}

	cmd := exec.Command("nvim", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return "", err
	}

	file, err := os.Open(f.Name())
	bytes, err := ioutil.ReadAll(file)
	file.Close()

	return string(bytes), err
}
