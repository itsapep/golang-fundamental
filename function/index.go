package main

import (
	"fmt"
)

func main() {
	// address := "Surabaya"
	// // store string value from the function
	// nama, alamat := printMessage("apep", address)
	// fmt.Print(nama, alamat)

	// province, region := location("Jakarta Selatan")
	// fmt.Printf("saya tinggal di %s dalam provinsi %s\n", province, region)
	// // the varialbe has already assigned, hence ":" is not required
	// province, region = location("Jakarta Timur")
	// fmt.Printf("saya tinggal di %s dalam provinsi %s\n", province, region)
	// province, region = location("Bandar Lampung")
	// fmt.Printf("saya tinggal di %s dalam provinsi %s\n", province, region)

	// // variadic
	// var animals =[]string("ayam","bebek","cicak")
	// enigmaZoo(animals)
	// fmt.Println()
	// enigmaZooVariadic(animals...) // cara pertama
	// fmt.Println()
	// enigmaZooVariadic("ayam","bebek","cicak") //cara kedua
	// nama := "Jution"
	// alamat := "Ragunan"
	// pendidikan := "S1"
	// skills := []string{"Ayam", "Bebek", "Cicak"}
	// portofolio(nama, alamat, pendidikan, skills...)
	// // using string.Join()
	// portofolio("Jution","Ragunan","S1","Tiduran","menghilang")

	// // closure
	// printGreetingFunc := greeting
	// fmt.Println(printGreetingFunc("Nico"))

	// // callback function
	// getValidEmail("juti@gmail.com", emailFilter)
	// getValidEmail("adalk/hdasdhasi", emailFilter)

	// anonymous function
	var greeting = func() {
		fmt.Println("Hai, aku anonim")
	}
	greeting()
}

// func printMessage(name, address string) {
// 	fmt.Printf("Hola, amigos! soy %s\n", name)
// 	fmt.Printf("vivo en %s\n", address)
// }

// // single return
// func printMessage(name, address string) string {
// 	// Sprintf only stores the string value into the variable "result"
// 	result := fmt.Sprintf("Hola, amigos! soy %s\nvivo en %s\n", name, address)
// 	return result
// }

// // double and more return
// func printMessage(name, address string) (string, string) {
// 	// Sprintf only stores the string value into the variable "result"
// 	getName := fmt.Sprintf("Hola, amigos! soy %s\n", name)
// 	getAddress := fmt.Sprintf("vivo en %s\n", address)
// 	return getName, getAddress
// }

// func location(city string) (province,region string /*string, string*/) {
// 	// declared as named result
// 	// var (
// 	// 	province, region string
// 	// )
// 	switch city {
// 	case "Jakarta Selatan", "Jakarta Barat":
// 		province, region = "DKI Jakarta", "Indonesia"
// 	case "Bandar Lampung", "Lampung Selatan":
// 		province, region = "Lampung", "Indonesia"
// 	default:
// 		province, region = "Tidak ada", "Tidak ada"
// 	}
// 	return // province, region
// }

// // breaking the function for certain condition
// func oddEven(number int) {
// 	if number == 0 {
// 		fmt.Println("Invalid number, can't xero number")
// 		return
// 	}
// 	if number%2 == 0 {
// 		fmt.Printf("Angka %d merupakan bilangan genap\n", number)
// 	} else {
// 		fmt.Printf("Angka %d merupakan bilangan ganjil\n", number)
// 	}
// }

// variadic function
/*
Variadic Function
-> Pembuatan function dengan parameter seperti biasa
-> Sama sendan membuat sebuah slice []string, []int, dll.
-> Penggunaannya cukup dengan (...tipeData), e.g. ...string
-> Mempunyai sebuah indeks
-> Penerapannya harus di akhir parameter yg dibuat
-> e.g. func numbers(number ...int) (v) | func numbers(number ...int, x, y, int) (x)
*/
// func enigmaZoo(animal []string) { // slice
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// }
// func enigmaZooVariadic(animal ...string) { // variadic
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// }
// func portofolio(nama, alamat, pendidikan string, skills ...string) {
// 	// var mySkill = strings.Join(skills,", ")
// 	fmt.Printf("Nama: %s\nAlamat: %s\nPendidikan: %s\nSkills: %s\n", nama, alamat, pendidikan, skills)
// }

// // function value (closure)
// /*
// sebuanh function yg bisa dijadikan sebuah value dalam sebua variable
// sehingga variabel tsb bertindak selayaknya sebuah function
// */
// func greeting(name string) string {
// 	return "salam kenal " + name
// }

// // callback function || function as parameter
// /*
// membuat sebuah callback function yang ditaruh pada parameter
// sehingga function bertindak sebagai parameter
// e.g. filter email, "@" -> a@gmail.com, y@gmail.com, dst.
// */
// func getValidEmail(email string, callback func(string) string) {
// 	emailFiltered := callback(email)
// 	fmt.Println("your email", emailFiltered)
// }
// func emailFilter(email string) string {
// 	if strings.Contains(email, "@") {
// 		return email
// 	} else {
// 		return "Not valid!"
// 	}
// }

// anonymous function
