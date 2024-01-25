package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationField(t *testing.T) {
	validate := validator.New()
	var user string

	// tidak bisa default value
	err := validate.Var(user, "required")

	if err != nil {
		t.Error(err.Error())
	}
}
