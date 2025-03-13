package config_test

import (
	"os"
	"testing"

	"github.com/lvyonghuan/Ubik-Util/config"
	"github.com/lvyonghuan/Ubik-Util/uerr"
)

func TestRead(t *testing.T) {
	// 创建test.json文件
	jsonContent := `{
		"string": "test",
		"number": 123,
		"boolean": true,
		"array": [1, 2, 3],
		"object": {
			"key": "value"
		},
		"null": null
	}`
	err := os.WriteFile("test.json", []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create example.json: %v", uerr.NewError(err))
	}
	defer os.Remove("test.json")

	// 定义接收结构体
	var configData struct {
		String  string                 `json:"string"`
		Number  int                    `json:"number"`
		Boolean bool                   `json:"boolean"`
		Array   []int                  `json:"array"`
		Object  map[string]interface{} `json:"object"`
		Null    interface{}            `json:"null"`
	}

	// 调用Read函数
	err = config.Read("test.json", &configData)
	if err != nil {
		t.Fatalf("Read function failed: %v", err)
	}

	// 验证读取结果
	if configData.String != "test" {
		t.Errorf("Expected string 'example', got %v", configData.String)
	}
	if configData.Number != 123 {
		t.Errorf("Expected number 123, got %v", configData.Number)
	}
	if configData.Boolean != true {
		t.Errorf("Expected boolean true, got %v", configData.Boolean)
	}
	if len(configData.Array) != 3 || configData.Array[0] != 1 || configData.Array[1] != 2 || configData.Array[2] != 3 {
		t.Errorf("Expected array [1, 2, 3], got %v", configData.Array)
	}
	if configData.Object["key"] != "value" {
		t.Errorf("Expected object key 'value', got %v", configData.Object["key"])
	}
	if configData.Null != nil {
		t.Errorf("Expected null, got %v", configData.Null)
	}
}
