package main

import (
	"fmt"
)

func main() {
	/*
		saya programmer golang
		saya programmer golang
		saya programmer golang
		saya programmer golang
		saya programmer golang
		saya programmer golang
	*/

	// basic for
	// for i := 0; i < 6; i++ {
	// 	fmt.Println(i, "- saya programmer golang")
	// }

	// for a while
	// i := 0
	// for i < 6 { // given the condition only
	// 	fmt.Println(i, "- saya programmer golang")
	// 	i++ // avoiding infinity loop
	// }

	// for ever
	// assign without any statement
	// i := 0
	// j := 0
	// for {
	// 	fmt.Println(i, "- saya programmer golang")
	// 	i++

	// 	// to stop at certain condition
	// 	if i == 9 {
	// 		break
	// 	}
	// }

	// for {
	// 	fmt.Println(i, "- saya programmer golang")
	// 	i++
	// 	time.Sleep(1 * time.Second)
	// 	for {
	// 		fmt.Println(i, "- saya programmer golang")
	// 		i++
	// 		time.Sleep(1 * time.Second)
	// 		if j == 4 {
	// 			j = 0
	// 			break
	// 		}
	// 	}
	// 	if i == 4 {
	// 		break
	// 	}
	// }

	// break? force stop a process
	// continue? force process to continue
	// return? force stop out of the function, here it is main function

	// print even odd
	// 1, 3, 5, 7, dst. (continue)
	// 2, 4, 6, 8, dst. (break)
	for i := 1; i < 12; i++ {
		if i%2 == 1 {
			continue // this will force skip all the odd number to the even one
		}
		if i > 8 {
			break
		}
		fmt.Println("Number: ", i)
	}

}
