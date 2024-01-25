package test

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestStructValidationFail(t *testing.T) {
	loginRequest := LoginRequest{
		Username: "kiki",
		Password: "kiki",
	}

	validator := validator.New()
	err := validator.Struct(loginRequest)

	assert.NotNil(t, &err)
	fmt.Println(err.Error())
}

func TestStructValidationSuccess(t *testing.T) {
	loginRequest := LoginRequest{
		Username: "kiki@mail.com",
		Password: "kiki12345",
	}

	validator := validator.New()
	err := validator.Struct(loginRequest)

	assert.Nil(t, err)
}

func TestIterateValidationError(t *testing.T) {
	loginRequest := LoginRequest{
		Username: "kiki",
		Password: "kiki",
	}

	validate := validator.New()
	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error " + fieldError.Field() + " on tag " + fieldError.Tag() + " with error " + fieldError.Error() + "\n")
		}
	}
}

type UserRegisterRequest struct {
	Email           string `validate:"email,required"`
	Password        string `validate:"required,min=3,max=50"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

func TestStructCrossField(t *testing.T) {
	registerRequest := UserRegisterRequest{
		Email:           "abc@gmail.com",
		Password:        "abababab",
		ConfirmPassword: "abababab0",
	}

	validate := validator.New()
	err := validate.Struct(registerRequest)
	assert.NotNil(t, err)
	fmt.Println(err)
}

type Address struct {
	City    string `validate:"required"`
	Country string `validate:"required"`
}

type User struct {
	Id      int     `validate:"required"`
	Name    string  `validate:"required"`
	Address Address `validate:"required"`
}

func TestValidateNestedStruct(t *testing.T) {
	createUserReq := User{
		Id:   1,
		Name: "A",
		Address: Address{
			City:    "",
			Country: "B",
		},
	}

	validate := validator.New()
	err := validate.Struct(createUserReq)
	assert.NotNil(t, err)
	fmt.Println(err)
}
