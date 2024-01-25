package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println("Channel kembalian: ", time)
}

func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)

	timerCh := <-channel // butuh channelnya saja
	fmt.Println("Channel kembalian: ", timerCh)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() { // nunggu lima detik dulu
		fmt.Println("Time inside func() ", time.Now()) // baru ini dieksekusi
		group.Done()
	})

	fmt.Println("Time outside func() ", time.Now())
	group.Wait()
}
