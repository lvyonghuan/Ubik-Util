package uerr_test

import (
	"errors"
	"fmt"
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

func TestExtractErrorReturnsNilForNil(t *testing.T) {
	if got := uerr.ExtractError(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestExtractErrorReturnsPlainErrorUnchanged(t *testing.T) {
	plain := errors.New("plain error")
	got := uerr.ExtractError(plain)
	if got == nil {
		t.Fatal("expected non-nil error")
	}
	if got.Error() != "plain error" {
		t.Fatalf("expected %q, got %q", "plain error", got.Error())
	}
	var ubikError uerr.UbikError
	if errors.As(got, &ubikError) {
		t.Fatal("expected plain error not to be UbikError")
	}
}

func TestExtractErrorReturnsUbikErrorWhenGivenUbikErrorValue(t *testing.T) {
	ub := uerr.NewError(errors.New("boom"))
	var err error = ub
	got := uerr.ExtractError(err)
	if got == nil {
		t.Fatal("expected non-nil error")
	}
	var ue uerr.UbikError
	ok := errors.As(got, &ue)
	if !ok {
		t.Fatalf("expected UbikError type, got %T", got)
	}
	if ue.MetaError().Error() != "boom" {
		t.Fatalf("expected inner error %q, got %q", "boom", ue.MetaError().Error())
	}
}

func TestExtractErrorUnwrapsWrappedUbikError(t *testing.T) {
	ub := uerr.NewError(errors.New("wrapped"))
	wrapped := fmt.Errorf("wrapper: %w", ub)
	got := uerr.ExtractError(wrapped)
	var ue uerr.UbikError
	ok := errors.As(got, &ue)
	if !ok {
		t.Fatalf("expected UbikError type, got %T", got)
	}
	if ue.MetaError().Error() != "wrapped" {
		t.Fatalf("expected inner error %q, got %q", "wrapped", ue.MetaError().Error())
	}
}
