package test

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[T]) ChangeFirst(newFirst T) {
	d.First = newFirst
}

// kalau tidak butuh field dari struct, maka kasih _ aja
func (d *Data[_]) SayHello(name string) {
	fmt.Println("Hello, " + name)
}

func TestGenericStruct(t *testing.T) {
	data := Data[int]{
		First:  2,
		Second: 1,
	}
	data.ChangeFirst(1)
	fmt.Println(data)
	data.SayHello("AAA")
}
