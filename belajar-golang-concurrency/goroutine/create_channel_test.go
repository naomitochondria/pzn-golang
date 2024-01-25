package goroutine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	fmt.Println("Membuat channel")
	channel := make(chan string)
	defer close(channel)

	go func() {
		// time.Sleep(2 * time.Second)
		channel <- "Ini data channel" // Apabila channel tidak diisi maka program akan tetap menunggu
		fmt.Println("Selesai mengirimkan data ke channel")
	}()

	fmt.Println("Menunggu channel menerima data")
	data := <-channel
	fmt.Println("Channel telah menerima data", data)

	// time.Sleep(2 * time.Second)
}

func GenerateRandomNumber(channel chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Mulai menghasilkan random number")

	randomNumber := rand.Intn(10000)
	channel <- randomNumber

	fmt.Println("Selesai menghasilkan random number")
}

func TestGenerateRandomNumber(t *testing.T) {
	channel := make(chan int)
	defer close(channel)

	go GenerateRandomNumber(channel)

	randomNumberResult := <-channel
	fmt.Println(randomNumberResult)

	fmt.Println("Program selesai")
	time.Sleep(2 * time.Second)
}
