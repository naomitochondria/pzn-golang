package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(T)
}

/*
Untuk bisa menjadi turunan GetterSetter,
MyData harus mengimplementasikan baik genericnya
dan method-methodnya
*/
type MyData[T any] struct {
	value T
}

func (d *MyData[T]) GetValue() T {
	return d.value
}

func (d *MyData[T]) SetValue(newValue T) {
	d.value = newValue
}

func ChangeValue[T any](param GetterSetter[T], value T) {
	param.SetValue(value)
	fmt.Println(param.GetValue())
}

func TestGenericInterface(t *testing.T) {
	data := MyData[int]{}
	ChangeValue(&data, 900)

	assert.Equal(t, 900, data.GetValue())
}
