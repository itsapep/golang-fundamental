package main

import "fmt"

func main() {
	var custName string = "James" //"custName" is camel case
	var custAge uint8 = 26
	var custStatus bool = true
	var custAddress string = "tuban"

	// var 2xxx // innapropriate varname
	// var xxx2 string // allowed varname
	// var _xxx string // allowed varname

	/*
		%s for string
		%d for int
		%f for float
		%t for boolean
		%v for anything
	*/

	fmt.Printf("nama %s umur %d alamat %s status %v \n", custName, custAge, custAddress, custStatus)
	fmt.Println("halo")

	var (
		firstName, lastName string
		age                 int
	)

	firstName = "jono"
	lastName = "jini"
	age = 22
	fmt.Println(firstName, lastName, age)

	// multiple init
	var (
		bootcampName, bootcampAddress = "enigmacamp", "jaksel"
	)
	fmt.Println(bootcampName, bootcampAddress)

	// inferred type
	days, month, year := "monday", "may", 2022
	fmt.Println(days, month, year)

	// set a constant
	const phi = 3.14
	fmt.Println(phi)
	const (
		statusOK      = 200
		statusCreated = 201
	)
	fmt.Println(statusOK, statusCreated)

	// imutable and mutable
	officeName, _ := "PT enigma", "bootcamp"
	fmt.Println(officeName)

	// a := "harus isi"
	var b float32
	fmt.Print("nilai b: ", b) // zero value
}
