package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var onceCounter = 0

func OnlyOnce() {
	fmt.Println("OnlyOnce called!")
	onceCounter += 1
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		fmt.Println("Begin waitgroup...")
		group.Add(1)
		go func() {
			once.Do(OnlyOnce) // hanya dijalankan sekali walaupun masuk ke 100 goroutine
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("onceConter = ", onceCounter)
}
