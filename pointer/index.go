package main

import "fmt"

func main() {
	// variable biasa
	var num int
	fmt.Println(num)

	// variable dengan pointer
	var ptr *int
	fmt.Println(ptr)

	// mendapatkan alamat memori
	fmt.Println(&num)
	fmt.Println(&ptr)

}
