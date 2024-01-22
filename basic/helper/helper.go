package helper

import "fmt"

/*
* digolang jika nama nya diawali huruf besar maka akan bisa diakses oleh package lain
* jika nama nya diawali dengan huruf kecil maka tidak bisa diakses oleh package lain
 */

// access modifier => private (awal huruf kecil)
// var version = "1.0"
// access modifier => public (awal huruf besar)
// var Application = "golang"

func SayHello(name string) string {
	return "Hello " + name
}

// package initialization
// menggunakan func init()
func init() {
	fmt.Println("Init Helper")
}
