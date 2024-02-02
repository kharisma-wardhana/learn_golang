package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic digunakan untuk menggunakan data primitice secara aman dalam process concurrent
func TestAtomic(t *testing.T) {
	var group = sync.WaitGroup{}
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				// akan terjadi race condition
				// jika langsung pake data primitive
				// counter++
				// gunakan atomic untuk handle race condtion pada data primitive
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
