package ujson

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/lvyonghuan/Ubik-Util/uerr"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, uerr.NewError(err)
	}

	return jsonData, nil
}

func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return uerr.NewError(err)
	}

	return nil
}
