package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidaton(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil!")
	}
}
