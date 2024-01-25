package goroutine

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//

func TestCreateContext(t *testing.T) {
	// digunakan pada saat pertama membuat context secara manual
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

//

func TestContextWithValue(t *testing.T) {
	background := context.Background()
	contextA := context.WithValue(background, "A", "a")
	contextB := context.WithValue(background, "B", "b")

	contextC := context.WithValue(contextA, "C", "c")
	contextD := context.WithValue(contextA, "D", "d")

	contextE := context.WithValue(contextB, "E", "e")

	for _, ctx := range []context.Context{contextA, contextB,
		contextC, contextD, contextE} {
		fmt.Println(ctx)
	}

	fmt.Println("C -> C: ", contextC.Value("C")) // mendapatkan valuenya sendiri
	fmt.Println("E -> B: ", contextE.Value("B")) // bisa, karena kalau tidak ada naik ke parentnya
	fmt.Println("D -> B: ", contextD.Value("B")) // nill, karena tidak ada di context itu dan context parentnya
}

//

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

//

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

func TestContextWithTimenout(t *testing.T) {
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

//

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
