package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()                   // dijalankan no (2). Tidak bisa menerima return value dari fungsi
	fmt.Println("UpsUpsUpsUpsUpsUpsUps") // dijalankan no (3) karena strinynya besar
	fmt.Println("Ups")                   // dijalankan no (1) karena stringnya kecil

	time.Sleep(1 * time.Second)
}
