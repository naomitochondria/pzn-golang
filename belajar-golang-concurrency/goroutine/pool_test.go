package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool2(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("A")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {

			data := pool.Get()
			fmt.Println("Pool data =. ", data)

			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Pool now: ", pool.Get())
	fmt.Println("TestPool finished!")
}
