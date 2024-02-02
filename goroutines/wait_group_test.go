package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// WaitGroup digunakan untuk menunggu process goroutine hingga selesai dilakukan
// method Add(int) setelah process selesai bisa gunakan method Done()
// untuk menunggu semua process selesai method Wait()

func RunAsync(group *sync.WaitGroup) {
	// warn perlu di done untuk avoid deadlock
	defer group.Done()

	// running 1 process async
	group.Add(1)
	fmt.Println("Run Async")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}
	group.Wait()
	fmt.Println("All Process Completed")
}
