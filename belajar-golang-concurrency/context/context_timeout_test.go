package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterV3(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second)
				counter += 1
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Jumlah goroutine saat ini : ", runtime.NumGoroutine())

	background := context.Background()
	ctx, cancel := context.WithTimeout(background, 5*time.Second)
	defer cancel() // walaupun ada timeout, tetap harus di cancel. Best practice

	destination := CreateCounterV3(ctx)
	for d := range destination {
		fmt.Println("Counter : ", d)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Jumlah goroutine saat ini : ", runtime.NumGoroutine())
}
