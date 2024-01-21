package main

import "fmt"

func main() {
	/*
		// int8 (-128 hingga 127) int mulai dari minus
		// int16 (-32768 hingga 32767)
		// int32 (-2147483648 hingga 2147483647)
		// int64 (-9223372036854775808 hingga -9223372036854775807)

		// uint8 (0 hingga 255) unsigned int mulai dari 0
		// uint16 (0 hingga 65535)
		// uint32 (0 hingga 4294967295)
		// uint64 (0 hingga 18446744073709551615)

		// float32 (1.18 x 10^-38 hingga 3.4 x 10^38)
		// float64 (2.23 x 10^-308 hingga 1.80 x 10^308)

		// alias
		// byte -> unit8
		// rune -> int32
		// int -> int32
		// uint -> uint32
	*/

	var a int
	a = 1
	fmt.Println("Satu", a)

	b := 1
	fmt.Println("Satu", b)

	/*
		// declare multiple variable
	*/
	var (
		firstname = "Kharisma"
		lastname  = "Wardhana"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	/*
		// convert data type
	*/
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	value16 := int16(value32)
	posiiveValue16 := uint16(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)
	fmt.Println(posiiveValue16)

	var k = firstname[0]
	kString := string(k)
	fmt.Println("Char K: ", k)
	fmt.Println("String K: ", kString)

	/*
		// Type declaration (membuat ulang data type baru dari data type yang sudah ada)
		// biasa digunakan untuk membuat alias terhadap data type yang sudah ada,
		// dengan tujuan agar lebih mudah dimengerti
	*/
	type NoKTP string

	var eKTP NoKTP = "34709029340"
	fmt.Println(eKTP)
	fmt.Println(NoKTP("909090"))
}
