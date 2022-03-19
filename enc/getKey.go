package enc

import (
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/neelkarma/ramble/constants"
)

func GetKey() []byte {
	file, _ := os.Open(constants.KeyPath)
	hexkey, _ := ioutil.ReadAll(file)
	key, _ := hex.DecodeString(string(hexkey))
	return key
}
