package main

import (
	"fmt"
	"testing"
)

func TestBreakContinue(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if i == 4 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
	}
}
