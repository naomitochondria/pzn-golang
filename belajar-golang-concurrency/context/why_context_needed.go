package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter += 1
		}
	}()

	return destination
}

func TestGoroutineLeak(t *testing.T) {
	fmt.Println("Jumlah goroutine saat ini: ", runtime.NumGoroutine())

	destination := CreateCounter()
	for d := range destination {
		fmt.Println("Counter : ", d)
		if d == 10 {
			break // yang di break hanya loop disini. Goroutine tetap berjalan.
		}
	}

	fmt.Println("Jumlah goroutine saat ini: ", runtime.NumGoroutine())
}

// Context dibutuhkan untuk mengatur masa hidupnya goroutine

func CreateCounterV2(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter += 1
			}
		}
	}()

	return destination
}

func TestCancelContext(t *testing.T) {
	fmt.Println("Jumlah goroutine saat ini: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterV2(ctx)

	for d := range destination {
		fmt.Println("Counter = ", d)
		if d == 10 {
			break
		}
	}
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Jumlah goroutine saat ini: ", runtime.NumGoroutine())
}
