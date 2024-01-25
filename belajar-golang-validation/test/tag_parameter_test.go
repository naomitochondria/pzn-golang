package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestParameter(t *testing.T) {
	validator := validator.New()
	name := "Abcdefghijklmn"

	err := validator.Var(name, "required,alpha,min=3,max=10")
	if err != nil {
		t.Error(err.Error())
	}
}
