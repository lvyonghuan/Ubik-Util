package uconfig

import (
	"os"

	"github.com/lvyonghuan/Ubik-Util/uerr"
	"github.com/lvyonghuan/Ubik-Util/ujson"
)

// Read the configuration file
// Pass in the profile path with the reception struct pointer
// An error message is returned
func Read(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return uerr.NewError(err)
	}

	err = ujson.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
