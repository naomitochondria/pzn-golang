package test

import (
	"fmt"
	"go-testing/controller"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Dijalankan di semua test
func TestMain(t *testing.M) {
	// Before test
	fmt.Println("Before test...")

	t.Run()

	// After test
	fmt.Println("After test...")
}

func TestHelloWorldAsset1(t *testing.T) {
	stringTest := controller.HelloWorld("Cupa")
	stringAnswer := "Hello, Cupa"

	assert.Equal(t, stringAnswer, stringTest) // Assert = tanpa if, nyocokin stringAnswer-stringTest, dan manggil Fail()
	fmt.Println("Hello Word Assert 1 Done!")
}

func TestHelloWorldAsset2(t *testing.T) {
	stringTest := controller.HelloWorld("Cupa")
	stringAnswer := "Hello Cupa"

	require.Equal(t, stringAnswer, stringTest) // Require = tanpa if, nyocokin stringAnswer-stringTest, dan manggil FailNow()
	fmt.Println("Hello Word Assert 2 Done!")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Di linux tidak bisa jalan!") // Tetap PASS tapi di bawahnya di skip
	}
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}
