package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func InChannel(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Test..."
	fmt.Println("Selesai memberikan data ke channel!")
}

func OutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println("Data: ", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go InChannel(channel)
	go OutChannel(channel)

	time.Sleep(3 * time.Second)
}
