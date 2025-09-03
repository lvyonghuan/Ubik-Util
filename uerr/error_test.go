package uerr_test

import (
	"errors"
	"testing"

	"github.com/lvyonghuan/Ubik-Util/uerr"
)

func TestNewUbikError(t *testing.T) {
	err := uerr.NewError(errors.New("an error occurred"))
	if err.MetaError().Error() != "an error occurred" {
		t.Errorf("expected 'An error occurred', got %s", err.MetaError().Error())
	}

	t.Log(err.Error())
}

func TestUbikErrorEqualError(t *testing.T) {
	ubikErr := uerr.NewError(errors.New("an error occurred"))
	var err error
	err = ubikErr
	if err.Error() != ubikErr.Error() {
		t.Errorf("expected %s, got %s", ubikErr.Error(), err.Error())
	} else {
		t.Log("success")
	}
}
