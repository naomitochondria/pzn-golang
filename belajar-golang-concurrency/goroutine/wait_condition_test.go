package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func WaitCondition(cond *sync.Cond, value int, group *sync.WaitGroup) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait() // setelah saya locking saya boleh lanjut tidak ke proses selanjutnya?
	fmt.Println("Done ", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	var locker = sync.Mutex{}
	var cond = sync.NewCond(&locker)
	var group = sync.WaitGroup{}
	var numGo = 10

	for i := 0; i < numGo; i++ {
		group.Add(1)
		go WaitCondition(cond, i, &group)
	}

	go func() {
		for i := 0; i < numGo; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // mensignal semua goroutine
	// }()

	group.Wait()
}
