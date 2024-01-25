package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println(x)
}

func TestMutexRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println(x)
}
