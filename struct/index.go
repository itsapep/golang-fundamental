package main

import (
	"fmt"
	"strings"
)

// create a struct
// private access modifier (only accessible within one package)
// created by using lower-case struct name
type student struct {
	nim, name string
	grade     int
	isDone    bool
	inst      institute // embedded struct
	courses   []course
}

type institute struct {
	instituteName, instituteAddress, instituteAccreditation string
	instituteSince                                          int
}
type highSchoolStudent struct {
	nis, name string
	grade     int
	institute // embedded struct, syntatic sugar
	courses   []course
}

type course struct {
	name string
}

func main() {
	// // create an object variable from struct
	var std student
	fmt.Println(std)

	// std.nim = "02311740000033"
	// std.name = "Yurham Afif"
	// std.grade = 100
	// std.isDone = true
	// fmt.Println("std: ", std)
	// fmt.Println("std.name: ", std.name)

	// // using literal struct
	// // data type MUST suit the assigned struct data type
	// var std2 = student{"3927273921", "Nicho", 1020, true}
	// fmt.Println("std2: ", std2)
	// fmt.Println("std2.name: ", std2.name)

	// std3 := student{}
	// fmt.Println("std3: ", std3)

	// // struct declaration without type alias
	// // assignment value directly inserted
	// std4 := struct {
	// 	nama string
	// }{
	// 	"jono",
	// }
	// fmt.Println("std4: ", std4)

	// // struct declaration without type alias
	// var std5 struct {
	// 	alamat string
	// }
	// std5.alamat = "ragunan"
	// fmt.Println("std5: ", std5.alamat)

	// // pointer application in struct
	// // taking reference from std2
	// // https://stackoverflow.com/questions/60469860/why-do-you-have-to-use-an-asterisk-for-pointers-but-not-struct-pointers
	// var stdPointer *student = &std2
	// fmt.Printf("std2 (address): %p\n", &std2)
	// fmt.Printf("stdPointer (address): %p\n", stdPointer)
	// fmt.Println("stdPointer.name: ", stdPointer.name)
	// // accessing value which is not a pointer data type because prop name: string
	// // fmt.Printf("*stdPointer.name: %p\n",*stdPointer.name)
	// fmt.Printf("*stdPointer.name: %v\n", (*stdPointer).name)

	// // create new variable using keyword new() in struct
	// // let's use student struct
	// var std3New = new(student)
	// // this will define the struct object
	// fmt.Println("std3New: ", std3New) // &{    0 false}
	// fmt.Printf("std3New (address): %p\n", &std3New)

	// // create slice,maps,channels using keyword make()

	xyzStudent := []student{
		{
			nim:  "8234632",
			name: "niwxnisx",
			inst: institute{
				instituteName:          "polinema",
				instituteAddress:       "malang",
				instituteAccreditation: "A",
				instituteSince:         1962,
			},
			courses: []course{
				{
					name: "matematika diskrit",
				}, {
					name: "algoritma dan struktur data",
				}, {
					name: "dasar pemrograman",
				},
			},
		},
		{
			nim:  "731923921839",
			name: "uhfwsna xalnxak",
			inst: institute{
				instituteName:          "polinema",
				instituteAddress:       "malang",
				instituteAccreditation: "A",
				instituteSince:         1962,
			},
			courses: []course{
				{
					name: "matematika diskrit",
				}, {
					name: "algoritma dan struktur data",
				}, {
					name: "dasar pemrograman",
				},
			},
		},
	}
	// fmt.Println(xyzStudent)
	// fmt.Println()
	// fmt.Printf("%+v \n", xyzStudent)

	// xyzStudent = append(xyzStudent, student{
	// 	nim:  "123456",
	// 	name: "uqusqwjnbiqw qwertyui",
	// 	inst: institute{
	// 		instituteName:          "polinema",
	// 		instituteAddress:       "malang",
	// 		instituteAccreditation: "A",
	// 		instituteSince:         1962,
	// 	},
	// 	courses: []course{
	// 		{
	// 			name: "matematika diskrit",
	// 		}, {
	// 			name: "algoritma dan struktur data",
	// 		}, {
	// 			name: "dasar pemrograman",
	// 		},
	// 	},
	// })

	addStudent(&xyzStudent, student{
		nim:  "123456",
		name: "uqusqwjnbiqw qwertyui",
		inst: institute{
			instituteName:          "polinema",
			instituteAddress:       "malang",
			instituteAccreditation: "A",
			instituteSince:         1962,
		},
		courses: []course{
			{
				name: "matematika diskrit",
			}, {
				name: "algoritma dan struktur data",
			}, {
				name: "dasar pemrograman",
			},
		},
	})

	// for _, std := range xyzStudent {
	// 	fmt.Println(std.name, std.nim, std.inst.instituteName)
	// 	for _, course := range std.courses {
	// 		fmt.Println(course)
	// 	}
	// 	fmt.Println(strings.Repeat("-", 50))
	// }

	printDatabase(xyzStudent)
	deleteDatabase(&xyzStudent, 1)
	printDatabase(xyzStudent)
	printExactDatabase(xyzStudent, 1)
}

func addStudent(data *[]student, std student) {
	*data = append(*data, std)
}

func printDatabase(data []student) {
	for _, std := range data {
		fmt.Println(std.name, std.nim, std.inst.instituteName)
		for _, course := range std.courses {
			fmt.Println(course)
		}
		fmt.Println(strings.Repeat("-", 50))
	}
}

func deleteDatabase(data *[]student, index int) {
	// a := []string{"A", "B", "C", "D", "E"}
	// i := 2

	// Remove the element at index i from a.
	// copy(data[index:], data[index+1:]) // Shift a[i+1:] left one index.
	// data[len(data)-1] = {  0 false {   0} []}    // Erase last element (write zero value).
	// a = a[:len(a)-1]     // Truncate slice.
	*data = append((*data)[0:index], (*data)[index+1:]...)
	// fmt.Println(data) // [A B D E]
}

func printExactDatabase(data []student, index int) {
	index--
	std := data[index]
	fmt.Println(std.name, std.nim, std.inst.instituteName)
	for _, course := range std.courses {
		fmt.Println(course)
	}
	fmt.Println(strings.Repeat("-", 50))
}
