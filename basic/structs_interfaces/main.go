package main

import (
	"fmt"
	"strconv"
)

// struct dalam golang mirip dengan class di java
// struct template untuk data atau prototype data
// bisa membuat object dari struct
// nama field dan nama struct digunakan PascalCase (huruf besar diawal)
// method adalah function yang menempel pada struct
type Customer struct {
	Name, Address string // field
	Age           int
}

// interface type data abstract
// interface sebagai contract biasa nya diimplementasikan ke struct
// golang bukan pemrograman berorientasi object
type HasName interface {
	GetName() string //method
}

type HasYearOfBirth interface {
	GetYearOfBirth(age int) string
}

// interface kosong alias any
// sehingga dapat digunakan untuk semua type data (sama dengan java.lang.Object)
// type any = interface{}
func ups() any {
	return 1
}

// secara default digolang tidak ada null melainkan default value (string "", int 0, bool false)
// nil hanya bisa digunakan pada interface, function, map, slice, pointer, dan channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{"name": name}
}

// cara implementasi interface digolang beda dengan dijava
// cukup dengan menyamakan contract saja
func (customer Customer) GetName() string { //customer sudah implement interface hasName
	return customer.Name
}

func (customer Customer) GetYearOfBirth(age int) string {
	return strconv.Itoa(customer.Age)
}

func main() {
	var kharis Customer
	fmt.Println("default value struct:", kharis)
	kharis.Name = "Kharisma Wardhana"
	kharis.Address = "Jogja"
	kharis.Age = 30
	fmt.Println(kharis)
	fmt.Println(kharis.Name)
	fmt.Println(kharis.Address)
	fmt.Println(kharis.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     28,
	}
	fmt.Println(joko)

	koko := Customer{"Koko", "Semarang", 30} //urutan harus sesuai
	fmt.Println(koko)

	// memanggil method pada struct Customer
	kharis.sayHello(joko.Name)
	koko.sayHello(joko.Name)

	sayHelloInterface(kharis) // memanggil method GetName dari Customer(yg sudah implement HasName)
	getYear(kharis)
}

// sample method yang ada pada struct Customer
// untuk memanggil method ini harus dari object Customer
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "nama saya", customer.Name)
}

func sayHelloInterface(val HasName) {
	fmt.Println(val.GetName())
}

func getYear(customer Customer) {
	fmt.Println(customer.GetYearOfBirth(customer.Age))
}
