package main

import (
	"fmt"
)

// function type declaration
type Filter func(string) string

type Blacklist func(string) bool

func main() {
	fmt.Println("Running function main")

	name := "kharis"
	if name == "kharis" {
		fmt.Println(name)
	} else {
		fmt.Println("Nama bukan kharis")
	}

	/*
		// contoh condition dengan short statement
		// length := len(name)
		// if length > 3 {}
		// dengan short statement maka penulisan berubah menjadi
		// if length := len(name); length > 3 {}
	*/
	if length := len(name); length < 3 {
		fmt.Println("nama pendek")
	} else {
		fmt.Println("nama benar")
	}

	switch name {
	case "Aris":
		fmt.Println("hello aris")
	case "Joko":
		fmt.Println("dimana aris?")
	default:
		fmt.Println("kenalan dong?")
	}

	/*
		// switch dengan short statement
	*/
	switch length := len(name); length < 3 {
	case true:
		fmt.Println("nama terlalu pendek")
	case false:
		fmt.Println("nama benar")
	}

	/*
		// switch tanpa condition
		// disarankan menggunakan if else
	*/
	length := len(name)
	switch {
	case length < 2:
		fmt.Println("nama terlalu pendek")
	case length > 10:
		fmt.Println("nama terlalu panjang")
	}

	/*
		// looping di golan
		// counter := 1
		// for counter <= 10 {
		// 	fmt.Println("Looping ke ", counter)
		// 	counter++
		// }
		// dalam for bisa ditambahkan 2 statement
		// for initStatement; statement; postStatement
		// for range untuk iteration data collection (array, slice, map)
	*/
	fmt.Println("looping dengan statement 10 data dengan continue index 2 dan break index ke 8")
	for count := 1; count <= 10; count++ {
		if count == 2 {
			continue
		}
		fmt.Println("looping dengan statement ke ", count)
		if count == 8 {
			break
		}
	}

	days := []string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}
	for i, day := range days { // jika index tidak dibutuhkan maka bisa diganti dengan (_) jadi for _,day := range days
		fmt.Println("index ", i, " day: ", day)
	}

	/*
		// function = first class citizen
		// function dianggap type data digolang dan bisa disimpan dalam var
	*/

	// memanggil function dengan parameter
	sayHello("Kharisma", "Wardhana")

	// memanggil function dengan ignore return value
	firstname, _ := getFullname()
	fmt.Println(firstname)

	// memanggil function dengan naming return value
	fmt.Println(getCompletedName())

	// memanggil function dengan parameter
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(getOddNumber(numbers))

	// memanggil variadic function
	fmt.Println("total", sumAll(1, 2, 3, 4, 5))
	// menggunakan slice dalam variadic function dengan spread operator
	fmt.Println("total", sumAll(numbers...))

	// memanggil function dari var (function as a value)
	goodbye := getGoodbye
	fmt.Println(goodbye("Kharis"))

	// function as parameter
	sayHelloWithFilter(firstname, spamFilter)

	// function type declaration
	sayHelloWithAlias(firstname, spamFilter)

	// anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser(firstname, blacklist)
	registerUser("anjing", func(name string) bool { return name == "anjing" })

	// recursive function (function yang memanggil dirinya sendiri)
	fmt.Println(factorialLoop(5))

	runApp(false)
	// defer ketika panic
	runApp(true)
}

/*
// sample membuat function di golang
// func namaFunction(namaParameter typeParameter, namaParameter2 typeParameter2) typeReturn
*/
func sayHello(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func getFullname() (string, string) {
	return "Kharisma", "Wardhana"
}

func getCompletedName() (firstname, middlename, lastname string) {
	firstname = "kharisma"
	middlename = "nanda"
	lastname = "wardhana"
	return firstname, middlename, lastname
}

func getOddNumber(number []int) []int {
	result := make([]int, len(number))
	for i, num := range number {
		if num%2 == 0 {
			result[i] = num
		}
	}
	return result
}

// sample variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func getGoodbye(name string) string {
	return "Bye " + name
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func sayHelloWithAlias(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func factorialLoop(val int) int {
	if val == 1 {
		return 1
	}
	return val * factorialLoop(val-1)
}

func logging() {
	fmt.Println("Selesai memanggil function")
	// recover harus digunakan dalam defer function
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error", message)
	}
}

func runApp(error bool) {
	// defer akan dipanggil ketika akhir function
	defer logging()

	// panic digunakan untuk menghentikan program namun defer akan tetap dieksekusi
	if error {
		panic("Error")
	}

	fmt.Println("Selesai Run App")
}
