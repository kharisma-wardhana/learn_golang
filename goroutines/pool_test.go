package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool digunakan untuk menyimpan data yang selanjutnya digunakan dan dikembalikan kedalam Pool
// untuk manage connection ke database
// Pool sudah aman dari race condition
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Kharis")
	pool.Put("Nanda")
	pool.Put("Wardhana")

	for i := 0; i < 10; i++ {
		go func() {
			// warn urutan tidak sesuai
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	fmt.Println(pool.Get())
	time.Sleep(11 * time.Second)
}
