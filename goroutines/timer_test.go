package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
	fmt.Println(<-channel)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	// digunakan untuk delay job
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println("DONE", time.Now())
	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		time.Sleep(100 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(5 * time.Second)
	for tick := range channel {
		fmt.Println(tick)
	}
}

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	// untuk merubah jumlah thread
	// biasa nya jumlah thread = total cpu core
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", totalGoroutine)

	group.Wait()
}
