package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// handle race condition
// perubahan data terhadap 1 var yang disharing dalam beberapa goroutine akan ada kemungkinan kena race condition
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; 1 < 1000; i++ {
		// 1000 goroutine
		// ketika ada 2 goroutine yang akses data x secara bersamaan
		// maka value x akan sama
		go func() {
			for j := 1; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	// expected result 100000
	fmt.Println("Counter", x)
}

// Mutex (Mutual Exclusion): untuk locking code program, butuh diunlock
// dengan menggunakan mutex hanya ada 1 goroutine yang diperbolehkan untuk melakukan process
func TestMutexRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; 1 < 1000; i++ {
		// 1000 goroutine
		// dengan locking tiap 1 goroutine
		go func() {
			for j := 1; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

// RWNutex (Read Write Mutex)
// digunakan untuk lock process Read dan process Write
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}
