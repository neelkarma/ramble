package entries

import (
	"io/ioutil"
	"time"

	"github.com/neelkarma/ramble/constants"
)

func GetAll() map[string]string {
	files, _ := ioutil.ReadDir(constants.EntriesPath)
	filemap := make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		entryTime, err := time.Parse(constants.EntryFileTimeFormat, file.Name())
		if err != nil {
			panic(err)
		}
		filemap[entryTime.Format(constants.EntryDisplayTimeFormat)] = file.Name()
	}
	return filemap
}
