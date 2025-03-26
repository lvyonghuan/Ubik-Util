package ujson_test

import (
	"testing"

	"github.com/lvyonghuan/Ubik-Util/uerr"
	"github.com/lvyonghuan/Ubik-Util/ujson"
	"github.com/stretchr/testify/assert"
)

func TestMarshalValidInput(t *testing.T) {
	input := map[string]string{"key": "value"}
	expected := `{"key":"value"}`

	result, err := ujson.Marshal(input)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(result))
}

func TestMarshalInvalidInput(t *testing.T) {
	input := make(chan int)

	result, err := ujson.Marshal(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.IsType(t, uerr.UbikError{}, err)
}

func TestUnmarshalValidInput(t *testing.T) {
	data := []byte(`{"key":"value"}`)
	var result map[string]string

	err := ujson.Unmarshal(data, &result)

	assert.NoError(t, err)
	assert.Equal(t, "value", result["key"])
}

func TestUnmarshalInvalidInput(t *testing.T) {
	data := []byte(`{"key":value}`)
	var result map[string]string

	err := ujson.Unmarshal(data, &result)

	assert.Error(t, err)
	assert.IsType(t, uerr.UbikError{}, err)
}
