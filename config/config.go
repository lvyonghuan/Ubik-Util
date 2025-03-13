package config

import (
	"Ubik-Util/uerr"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Read 读取配置文件
// 传入配置文件路径与接收结构体指针
// 返回错误信息
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
