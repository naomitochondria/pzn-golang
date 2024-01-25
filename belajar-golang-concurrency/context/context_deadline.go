package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterV4(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Goroutine saat ini : ", runtime.NumGoroutine())

	background := context.Background()
	ctx, cancel := context.WithDeadline(background, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterV4(ctx)
	for d := range destination {
		fmt.Println("Counter : ", d)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine saat ini : ", runtime.NumGoroutine())
}
