package main

import "fmt"

func main() {
	// is it odd or even?
	// a := 20
	// if a%2 == 0 {
	// 	println("this is even number")
	// } else {
	// 	fmt.Println("this is odd number")
	// }

	// =============================================

	// GPA score calculation simulator
	// gpa := 3.5
	// if gpa >= 3.5 && gpa <= 4.0 {
	// 	fmt.Printf("GPA %.1f: Cumlaude",gpa)
	// }else if gpa>=3.0&&gpa<3.5{
	// 	fmt.Printf("GPA %.1f: memuaskan",gpa)
	// }else if gpa>=2.75&&gpa<3.0{
	// 	fmt.Printf("GPA %.1f: cukup",gpa)
	// }else{
	// 	fmt.Printf("GPA %.1f: kurang memuaskan",gpa)
	// }

	// =============================================

	// point := 8000.00
	// percent := point / 100

	// if percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad", percent, "%")
	// }

	// =============================================

	point := 8000.00

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%% perfect", percent)
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad", percent, "%")
	}
}
