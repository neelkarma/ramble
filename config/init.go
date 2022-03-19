package config

import (
	"encoding/hex"

	"github.com/neelkarma/ramble/enc"
	"github.com/neelkarma/ramble/passhash"
)

func Init(passphrase string) {
	hash, _ := passhash.HashPassphrase(passphrase)
	encIv := hex.EncodeToString(enc.GenerateIv())
	encSalt := hex.EncodeToString(enc.GenerateSalt())

	config := &Config{
		Passhash: hash,
		EncIv:    encIv,
		EncSalt:  encSalt,
	}

	WriteConfig(config)
}
