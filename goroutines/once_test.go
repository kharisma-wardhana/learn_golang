package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Once digunakan untuk memastikan func dieksekusi 1x saja
// jika ada multiple goroutine yg akses 1 func maka hanya yg pertama yg dieksekusi yg lain diskip

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai", counter)
}
