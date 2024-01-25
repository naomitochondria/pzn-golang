package test

import (
	"fmt"
	"go-testing/controller"
	"testing"
)

func TestGetGreeting1(t *testing.T) {
	greeting := controller.GetGreeting("Cupa", 10)
	if greeting != "Cekcek" {
		t.FailNow() // Berhenti saat itu juga
	}
	fmt.Println("TestGetGreeting1 done!") // tidak akan dijalankan
}

func TestGetGreeting2(t *testing.T) {
	greeting := controller.GetGreeting("Puma", 18)
	if greeting != "Testes" {
		t.Fail() // Berhenti nanti
	}
	fmt.Println("TestGetGreeting2 done!")
}

func TestGetGreeting3(t *testing.T) {
	greeting := controller.GetGreeting("Sisil", 10)
	if greeting != "Cekcek" {
		t.Error("ERROR TestGetGreeting3") // Ngeprint log dan panggil Fail()
	}
	fmt.Println("TestGetGreeting3 done!")
}

func TestGetGreeting4(t *testing.T) {
	greeting := controller.GetGreeting("Cassio", 10)
	if greeting != "Cekcek" {
		t.Fatal("ERROR TestGetGreeting4") // Ngeprint log dan panggil FailNow()
	}
	fmt.Println("TestGetGreeting4 done!")
}

/*
	go test -v ./... -run=TestGetGreeting
*/
