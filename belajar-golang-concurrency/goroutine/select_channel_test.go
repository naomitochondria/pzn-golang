package goroutine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateRandomNumber(channel chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Mulai menghasilkan random number")

	randomNumber := rand.Intn(10000)
	channel <- randomNumber

	fmt.Println("Selesai menghasilkan random number")
}

/*
Untuk nungguin kapan data dari channel keluar
Untuk mengambil data pada channel dengan benar
*/
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan int)
	channel2 := make(chan int)
	defer close(channel1)
	defer close(channel2)

	go generateRandomNumber(channel1)
	go generateRandomNumber(channel2)

	go func() {
		counter := 0
		for {
			if counter >= 2 {
				break
			}
			select {
			case data := <-channel1:
				fmt.Println("Data: ", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data: ", data)
				counter++
			default:
				fmt.Println("Menunggu data....") // dijalankan apabila channel belum berisi data
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
