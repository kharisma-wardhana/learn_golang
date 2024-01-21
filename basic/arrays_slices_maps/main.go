package main

import "fmt"

func main() {
	/*
		// Array di golang tidak bisa bertambah daya tampung nya
		// di golang tidak bisa hapus value array hanya bisa dikosongkan
	*/
	var names [3]string
	names[0] = "Kharisma"
	names[1] = "Nanda"
	names[2] = "Wardhana"
	fmt.Println(names)
	fmt.Println(names[2])

	values := [4]int{
		20,
		1,
		2024,
	}
	fmt.Println("Check length values: ", len(values))
	fmt.Println(values)
	fmt.Println("Sebelum diubah: ", values[2])
	values[2] = 10
	fmt.Println("Setelah diubah: ", values[2])

	/*
		// Opsi declare array lain, dengan spread symbol (...)
		// warn: opsi ini harus langsung dideclare value nya
	*/
	v := [...]int{
		0,
		10,
		100,
		1000,
	}
	fmt.Println("Check length v: ", len(v))
	fmt.Println("v: ", v)

	/*
		// Slice potongan dari data Array
		// ukuran slice bisa berubah secara dynamic
		// pointer: penunjuk data pertama di slice
		// length: panjang slice
		// capacity: kapasitas slice (data pointer ke akhir dari array)
		// notes length tidak boleh lebih dari capacity
		// v[0:2] -> pointer:0, length:2 , capacity:4
		// v[2:] -> pointer:2, length:3, capacity:3
	*/

	fmt.Println("slice v dari 0-1: ", v[0:2])
	fmt.Println("slice v dari 2-akhir: ", v[2:])
	fmt.Println("slice v dari awal-2: ", v[:3])
	fmt.Println("slice v dari awal-akhir: ", v[:])

	var slice []int = v[:3]
	fmt.Println("slice v: ", slice)
	fmt.Println("kapasitas slice: ", cap(slice))
	fmt.Println("length slice: ", len(slice))
	/*
		// append digunakan untuk menambahkan data ke posisi terakhir dari slice
		// jika kapasitasnya sudah full maka akan membuat array baru
	*/
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	fmt.Println("days: ", days)
	daysSlice := days[2:] //Rabu, Kamis, Jumat, Sabtu, Minggu
	fmt.Println("days[2:] => ", daysSlice)
	/*
		// ketika slice diupdate maka array juga akan terupdate
	*/
	daysSlice[0] = "Rabu_uwu"
	fmt.Println("updated daysSlice =>", daysSlice)
	fmt.Println("updated days =>", days)

	fmt.Println("length daysSlice => ", len(daysSlice))
	fmt.Println("kapasitas daysSlice => ", cap(daysSlice))
	/*
		// ketika append maka akan membuat array baru
		// sehingga array yang awal jd tidak terpengaruh dengan perubahan yg dilakukan
	*/
	daysSliceAppend := append(daysSlice, "Hari Baru")
	fmt.Println("check append => ", daysSliceAppend)
	daysSliceAppend[0] = "Update Hari"
	fmt.Println("days => ", days)

	/*
		// membuat slice dengan make(type array, len, cap)
	*/
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Kharisma"
	newSlice[1] = "Nanda"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice2 := append(newSlice, "Wardhana")
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
	newSlice2[0] = "GantiNama"
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println("from slice: ", fromSlice)
	fmt.Println("to slice: ", toSlice)

	/*
		// WARN:
		// harus hati" ketika membuat array atopun slice
		// dalam golang kebanyakan menggunakan slice
	*/
	iniArray := [...]int{1, 2, 3}
	iniArray2 := [4]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3}
	fmt.Println("iniArray: ", iniArray)
	fmt.Println("iniArray2: ", iniArray2)
	fmt.Println("iniSlice: ", iniSlice)

	/*
		// Map = pair key-value (bisa ditambahkan sebanyak-banyaknya dengan key yang berbeda)
		// ex: map[key]value {"name": "kharis"}
	*/
	person := map[string]string{
		"name": "kharisma",
		"age":  "30",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"]) // jika key tidak ketemu maka pake default value (string => "")

	fmt.Println("before: ", person["age"])
	person["age"] = "28"
	fmt.Println("after: ", person["age"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Author"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(book["ups"])
	delete(book, "ups")
	fmt.Println("after delete key 'ups'", book)
}
