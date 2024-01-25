package context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(t *testing.T) {
	// digunakan pada saat pertama membuat context secara manual
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

//

func TestContextWithValue(t *testing.T) {
	background := context.Background()
	contextA := context.WithValue(background, "A", "a")
	contextB := context.WithValue(background, "B", "b")

	contextC := context.WithValue(contextA, "C", "c")
	contextD := context.WithValue(contextA, "D", "d")

	contextE := context.WithValue(contextB, "E", "e")

	for _, ctx := range []context.Context{contextA, contextB,
		contextC, contextD, contextE} {
		fmt.Println(ctx)
	}

	fmt.Println("C -> C: ", contextC.Value("C")) // mendapatkan valuenya sendiri
	fmt.Println("E -> B: ", contextE.Value("B")) // bisa, karena kalau tidak ada naik ke parentnya
	fmt.Println("D -> B: ", contextD.Value("B")) // nill, karena tidak ada di context itu dan context parentnya
}
