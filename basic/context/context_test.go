package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()
	// cval akan jd child dari background
	cval := context.WithValue(background, "Name", "Kharis")
	cvalB := context.WithValue(cval, "Age", "30")

	fmt.Println(background)
	fmt.Println(todo)
	fmt.Println(cval)
	fmt.Println(cvalB)

	fmt.Println(cval.Value("Name"))
	fmt.Println(cvalB.Value("Name"))
	fmt.Println(cvalB.Value("Age"))
	fmt.Println(cvalB.Value("Address"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()
	return destination
}

func TestCreateCounter(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent)
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	// cancel() //mengirim sinyal cancel ke context, pada with cancel harus dipanggil manual

	// terjadi goroutine leak jika jumlah goroutine diawal != diakhir (default 2 tp ditemukan 3)
	// ada 1 goroutine yang tidak terpakai dan masih menyala
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println(runtime.NumGoroutine())
}
