package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map aman untuk concurrent (aman dari race condition)
// Store(key, value)
// Load(key)
// Delete(key)
// Range(function(key,value))
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}
	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

func AddToMap(data *sync.Map, val int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(val, val)
}
