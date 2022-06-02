package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var i int
	var err interface{}
	fmt.Println("Enter number of threads:")
	for {
		_, err = fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else {
			fmt.Println("Got: " + strconv.Itoa(i))
			break
		}
	}
	//Todo
}
