package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// number: int, uint, float
	var numOne int = -1
	var numTwo uint = 1
	var numThree float32 = 28

	fmt.Printf("numOne: %d \n", numOne)
	fmt.Printf("numTwo: %d \n", numTwo)
	fmt.Printf("numThree: %.2f \n", numThree)

	// boolean
	var isValid bool = true
	fmt.Printf("isValid: %t \n", isValid)

	// string
	var message string = "halo salam kenal"
	fmt.Printf("message: %s \n", message)

	// working with string
	// string literal
	getMessage := `hello.....
	hai
	gratzie
	hola
	`
	fmt.Println(getMessage)

	getData := `hari ini adalah hari jumat ke pasar membeli "apel, garam, gula"`
	// to avoid error if using doublequote, add backslash
	// getData:= "hari ini adalah hari jumat ke pasar membeli \"apel, garam, gula\""
	println(getData)

	// concatenate
	firstName := "yurham"
	lastName := "afif"
	// age := 22
	// fullName := firstName + " " + lastName + age // not applicable with different data types
	fullName := firstName + lastName
	println(fullName)

	println(strings.ToUpper(firstName))
	println(strings.ToLower(firstName))
	println(strings.Contains(firstName, "y")) // case sensitive
	println(strings.Count(lastName, "f"))     // case sensitive
	println(len(firstName))                   // amount of chars

	// convert data type
	var index int8 = 15
	// index = 128
	otherIndex := int16(index)
	// otherIndex = 32768
	println((otherIndex))

	var number int64 = 129
	otherNumber := int8(number)
	fmt.Println("number: ", number)           // 129
	fmt.Println("otherNumber: ", otherNumber) // -127

	// check data type
	fmt.Printf("tipe data index: %T\n", index)
	fmt.Printf("tipe data otherIndex: %T\n", otherIndex)

	var x int64 = 57
	y := float64(x)
	fmt.Printf("%d \n", x)
	fmt.Printf("%f", y)

	var f float32 = 80.99
	i := int32(f)
	fmt.Printf("%f \n", f)
	fmt.Printf("%d", i)

	// number to string
	var a int = 12
	b := fmt.Sprintf("%d", a)
	fmt.Printf("print with sprintf: %T \n", b)

	c := strconv.Itoa(a)
	fmt.Printf("print with strconv.itoa: %T \n", c)

	// string to number
	e := "12"
	g := "13"
	result := e + g
	fmt.Println(result)

	h, _ := strconv.Atoi(e)
	j, _ := strconv.Atoi(g)
	result2 := h + j
	fmt.Println(result2)

}
