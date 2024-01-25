package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (ba *BankAccount) AddBalance(n int) {
	ba.RWMutex.Lock()
	ba.Balance += n
	ba.RWMutex.Unlock()
}

func (ba *BankAccount) ReadBalance() int {
	var balance int
	ba.RWMutex.RLock()
	balance = ba.Balance
	ba.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	bankAccount := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			bankAccount.AddBalance(i)
			fmt.Println(bankAccount.ReadBalance())
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Final balance", bankAccount.Balance)
}
