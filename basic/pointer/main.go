package main

import "fmt"

/*
* secara default semua var di golang passing by value bukan reference
* jika mengirim sebuah var ke dalam function, method, ato var lain
* yg dikirim adalah duplikasi value nya
* ketika terjadi perubahan data nya maka data awalnya masih aman
 */

/*
 * Pointer : kemampuan membuat reference lokasi data di memory yang sama, tanpa duplikasi data yang sudah ada
 * di golang untuk pass by reference menggunakan pointer
 */

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Jogja", "DIY", "Indonesia"}
	// pass by value
	// copy value dari address1 ke address2
	address2 := address1
	address2.City = "Sleman"

	fmt.Println("address2", address2)
	fmt.Println("address1 pass by value:", address1)

	// pass by reference
	// address3 pointer ke address1 (address3 penunjuk ke address1)
	// jika ada perubahan maka data awal akan ikut berubah
	fmt.Println("address1", address1)

	// cara penulisan pointer
	// var address3 *Address = &address1 (pointer Address menunjuk ke alamat memory address1)
	address3 := &address1

	address3.City = "Bantul"
	fmt.Println("address3 (pointer ke address1)", address3)
	fmt.Println("address3 (memory address1)", &address3)
	fmt.Println("address3 (value address1)", *address3)
	fmt.Println("address1", address1)

	// untuk mengubah value address1
	// *address3 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println("address1", address1)

	// untuk menggubah penunjuk maka digunakan &
	address3 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address3)

	// var address4 *Address = new(Address)
	address4 := new(Address)
	address5 := address4
	address5.Country = "Indonesia"
	fmt.Println("address4", address4)
	fmt.Println("address5", address5)

	alamat := Address{"Jogja", "DIY", ""}
	// untuk pass by reference
	// menggunakan &address untuk menunjuk ke memory
	// sehingga yg dikirim ke function berupa reference bukan value
	changeAddressToIndonesia(&alamat)
	fmt.Println(alamat)

	kharis := Man{"Kharis"}
	kharis.Married()
	fmt.Println(kharis.Name)
}

// sample pointer di function
func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

// recomendation selalu menggunakan pointer di method
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}
