package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
RunAsynchronus
  - Membuat goroutine yang melakukan print Hello
*/
func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
