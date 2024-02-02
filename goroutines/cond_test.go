package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// cond untuk implementasi locking basis condition
// cond membutuhkan locker (Mutex / RWMutex) untuk implementasi lockingnya
// signal digunakan untuk memberitahu goroutine untuk tidak menunggu lagi
// broadcast untuk memberi
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(val int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	// warn ketika ada wait pastikan ada Signal()
	cond.Wait()

	fmt.Println("DONE", val)

	cond.L.Unlock()
}

func TestWaitCondition(t *testing.T) {
	for i := 0; i < 20; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 20; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
