package goroutine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Nama    string
	Balance uint64
}

func (ub *UserBalance) Lock() {
	ub.Mutex.Lock()
}

func (ub *UserBalance) Unlock() {
	ub.Mutex.Unlock()
}

func (ub *UserBalance) Change(amount uint64) {
	ub.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount uint64, i int) {
	fmt.Println("Transfer ke-" + strconv.Itoa(i))

	user1.Lock()
	fmt.Println("Lock user1 ", user1.Nama)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 ", user2.Nama)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	fmt.Println("Unlock user 1", user1.Nama)
	user1.Unlock()

	fmt.Println("Unlock user 2", user2.Nama)
	user2.Unlock()
}

func TestTransfer(t *testing.T) {
	user1 := UserBalance{
		Nama:    "A",
		Balance: 1000,
	}
	user2 := UserBalance{
		Nama:    "B",
		Balance: 1000,
	}

	// deadlock: saling menunggu
	go Transfer(&user1, &user2, 100, 1) // hasil: user1=900 user2=1100
	go Transfer(&user2, &user1, 300, 2) // hasil: user1=1200 user2=800
	// hasil deadlock: user1=900 user2=700

	time.Sleep(10 * time.Second)

	fmt.Println("User1 ", user1.Nama, " balance: ", user1.Balance)
	fmt.Println("User2 ", user2.Nama, " balance: ", user2.Balance)
}
