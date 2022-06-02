package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "mangga"
	fruits[1] = "anggur"
	fruits[2] = "jeruk"
	fruits[3] = "alpukat"

	// // cara memanggil/mencetak
	// fmt.Println(fruits)
	// fmt.Println()
	// fmt.Printf("%s %s %s %s\n", fruits[0], fruits[1], fruits[2], fruits[3])

	// %q untuk cetak tiap elemen
	fmt.Printf("%q", fruits)

	bootcamp := [2]string{"enigma", "bootcamp"}
	fmt.Println(bootcamp)

	/*
		inisiasi value:
		1. Horizontal -> bootcamp:=[2]string{"enigma","bootcamp"}
		2. Vertical -> bootcamp:=[2]string{
									"enigma",
									"bootcamp"}
	*/
	bootcampVertical := [2]string{
		"enigma",
		"bootcamp"}
	fmt.Println(bootcampVertical)
}
