package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestMultipleTag(t *testing.T) {
	validator := validator.New()
	user := "user123456"

	// kalau ada 1 yang tidak valid, maka langsung error
	// gagalnya di tag numeric
	err := validator.Var(user, "required,numeric")
	if err != nil {
		t.Error(err.Error())
	}
}
