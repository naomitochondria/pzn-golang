package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	/*
		- channel := make(chan string)
			maka akan blocking -> program berjalan terus sampai ada yang mengambil
		- channel := make(chan string, 1)
			tidak akan blocking karena masuk ke buffer
	*/
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "A"
		channel <- "B"
		channel <- "C"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	capacity := 10
	channel := make(chan string, capacity)

	go func() {
		for i := 0; i < capacity; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		defer close(channel)
	}()
	go func() { // kalau for di luar goroutine dan channel tidak di close maka program berjalan terus
		for data := range channel {
			fmt.Println(data)
		}
	}()

	time.Sleep(2 * time.Second)
}
