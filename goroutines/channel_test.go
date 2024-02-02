package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// sample
	// channel <- "Kharis"

	// sample pake goroutine
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Kharis"
		fmt.Println("Finish Send Data To Channel")
	}()

	// warn akan kena deadlock jika tidak ada data yg diterima
	data := <-channel
	fmt.Println("Get Channel From Var", data)
	fmt.Println("Get Data Channel", <-channel)
	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go SampleParameter(channel)

	data := <-channel
	fmt.Println("Get Channel From Var", data)
	fmt.Println("Get Data Channel", <-channel)
	time.Sleep(5 * time.Second)
}

func SampleParameter(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Kharis"
	fmt.Println("Finish Send Data To Channel")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Kharisma Wardhana"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// secara default channel hanya bisa menerima 1 data
// buffered channel
// buffer capacity (ex: 5 maka data ke-6 harus menunggu sampai ada buffer kosong)
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// tidak ada blocking ketika menggunakan buffer channel
	go func() {
		channel <- "Kharisma"
		fmt.Println("Length Channel 1:", len(channel))
		channel <- "Nanda"
		fmt.Println("Length Channel 2:", len(channel))
		channel <- "Wardhana"
	}()

	fmt.Println("Capacity Channel", cap(channel))
	fmt.Println("Length Channel 3:", len(channel))

	// membaca data buffer
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("Selesai")

	time.Sleep(2 * time.Second)
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Index %d", i)
		}
		// warn perlu di close agar tidak terkena deadlock
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(2 * time.Second)
}

// select channel
// untuk menggambil data dari multiple channel secara langsung (tanpa for range untuk tiap channel)
func TestSelectChannel(t *testing.T) {
	channelA := make(chan string)
	channelB := make(chan string)

	defer close(channelA)
	defer close(channelB)

	go SampleParameter(channelA)
	go SampleParameter(channelB)

	counter := 0
	for {
		select {
		case data := <-channelA:
			fmt.Println("data channel A", data)
			counter++
		case data := <-channelB:
			fmt.Println("data channel B", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// sample penggunaan default untuk menunggu data dari pengirim
// digunakan untuk melakukan sesuatu ketika menunggu data channel
func TestDefaultSelectChannel(t *testing.T) {
	channelA := make(chan string)
	channelB := make(chan string)

	defer close(channelA)
	defer close(channelB)

	go SampleParameter(channelA)
	go SampleParameter(channelB)

	counter := 0
	for {
		select {
		case data := <-channelA:
			fmt.Println("data channel A", data)
			counter++
		case data := <-channelB:
			fmt.Println("data channel B", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
