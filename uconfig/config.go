package uconfig

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/lvyonghuan/Ubik-Util/uerr"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Read the configuration file
// Pass in the profile path with the receive struct pointer
// An error message is returned
func Read(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return uerr.NewError(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return uerr.NewError(err)
	}

	return nil
}
