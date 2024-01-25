package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func StopTicker(ticker *time.Ticker) {
	fmt.Println("Stopping the ticker...")
	ticker.Stop()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // ticker itu interval
	maxIter := 5

	for time := range ticker.C {
		if maxIter == 0 {
			StopTicker(ticker)
		} else {
			fmt.Println(time)
			maxIter--
		}
	}
}

func TesTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
