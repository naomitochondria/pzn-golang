package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationTwoVariable(t *testing.T) {
	input := "password"
	realPassword := "password"

	validator := validator.New()
	err := validator.VarWithValue(input, realPassword, "eqfield")
	if err != nil {
		t.Error(err.Error())
	}
}
