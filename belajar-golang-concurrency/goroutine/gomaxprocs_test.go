package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	// total goroutine yang berjalan saat ini
	totalGoroutine := runtime.NumGoroutine() // default 2: 1 untuk menjalankan fungsinya, 1 untuk garbage collection
	fmt.Println(totalGoroutine)
}
