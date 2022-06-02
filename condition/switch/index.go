package main

import "fmt"

func main() {
	/*
		switch statement{
		case condition:
			// output
			break (x) // unnecessary
		default:
			// output
		}
	*/
	point := 100

	// switch point{
	// case 100:
	// 	fmt.Println("Perfect")
	// case 80:
	// 	fmt.Println("Good")
	// default:
	// 	fmt.Println("Not bad")
	// }

	// multiple case
	// switch point{
	// case 100,90,80:
	// 	fmt.Println("Perfect")
	// case 70,60:
	// 	fmt.Println("Good")
	// case 50:
	// 	fmt.Println("Good")
	// default:
	// 	fmt.Println("Not bad")
	// }

	// // switch like if else
	// switch {
	// case point == 100:
	// 	fmt.Println("Perfect")
	// case point >= 80 && point < 100:
	// 	fmt.Println("Good")
	// case point > 60 && point < 80:
	// 	fmt.Println("Good")
	// default:
	// 	fmt.Println("Not bad")
	// }

	// keyboard fallthrough: force to check 1 level under
	switch {
	case point == 100:
		fmt.Println("Perfect")
		fallthrough
	case point >= 80 && point < 100:
		fmt.Println("Good")
	case point > 60 && point < 80:
		fmt.Println("Good")
	default:
		fmt.Println("Not bad")
		fmt.Println("You need to learn more")
	}
}
