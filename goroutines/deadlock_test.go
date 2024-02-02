package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Deadlock
// keadaan process goroutine saling tunggu (biasanya karena process lock)
// simulasi deadlock
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) AddBalance(amount int) {
	user.Balance += amount
}

func (user *UserBalance) GetBalance() int {
	return user.Balance
}

func Transfer(userA *UserBalance, userB *UserBalance, amount int) {
	userA.Lock()
	fmt.Println("Lock UserA:", userA.Name)
	userA.AddBalance(-amount)

	time.Sleep(2 * time.Second)

	userB.Lock()
	fmt.Println("Lock UserB:", userB.Name)
	userB.AddBalance(amount)

	time.Sleep(2 * time.Second)

	userA.Unlock()
	userB.Unlock()
}

func PrintData(user *UserBalance) {
	fmt.Println("user:", user.Name, "balance:", user.Balance)
}

func TestDeadlock(t *testing.T) {
	userA := UserBalance{
		Name:    "Aries",
		Balance: 1000,
	}

	userB := UserBalance{
		Name:    "Nanda",
		Balance: 0,
	}

	PrintData(&userA)
	PrintData(&userB)

	// ada process locking di userA dan userB
	// goroutine akan saling menunggu
	go Transfer(&userA, &userB, 500)
	go Transfer(&userB, &userA, 600)

	time.Sleep(2 * time.Second)

	PrintData(&userA)
	PrintData(&userB)
}
