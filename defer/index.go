package main

import (
	"fmt"
)

func main() {
	/*
		Defer -> sebuah baris eksekusi yang terakhir
	*/
	// fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)

	// defer a()
	// b()
	// c()

	// number := 6
	// if number == 6 {
	// 	fmt.Println(4)
	// 	defer fmt.Println(6)
	// }
	// fmt.Println(5)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// func a() {
// 	fmt.Println("func a")
// }
// func b() {
// 	fmt.Println("func b")
// }
// func c() {
// 	fmt.Println("func c")
// }
