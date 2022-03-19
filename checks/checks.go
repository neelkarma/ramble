package checks

import (
	"errors"
	"os"

	"github.com/neelkarma/ramble/constants"
)

func DiaryExists() bool {
	_, err := os.Stat(constants.ConfigPath)
	return !errors.Is(err, os.ErrNotExist)
}

func DiaryUnlocked() bool {
	_, err := os.Stat(constants.KeyPath)
	return !errors.Is(err, os.ErrNotExist)
}

func TempFileExists() bool {
	_, err := os.Stat(constants.TempFilePath)
	return !errors.Is(err, os.ErrNotExist)
}
