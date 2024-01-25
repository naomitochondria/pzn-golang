package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToSyncMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	data.Store(value, value*10)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go AddToSyncMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ": ", value)
		return true
	})
}
