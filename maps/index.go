package main

import (
	"fmt"
	"sort"
)

func main() {
	// pembuatan map dengan keyword make
	mapScore := make(map[string]int)
	mapScore["Angga"] = 75
	mapScore["Sinta"] = 60
	mapScore["Roi"] = 85
	mapScore["Roi"] = 90 // assign new value

	fmt.Println()

	// pembuatan map
	var month = map[int]string{
		1: "Januari",
		2: "Februar",
		3: "Maret",
		4: "April",
		5: "Mei",
	}
	// cetak dengan for loop
	// this will show unsorted map
	for key, value := range month {
		fmt.Println(key, ":\t", value)
	}

	fmt.Println()

	// sort map
	keys := make([]int, 0, len(month))
	for k := range month {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, month[k])
	}
	fmt.Println()

	// we may delete the element(s) of map using 'key'
	// the command is delete()
	fmt.Println("before deleted: ", len(month))
	delete(month, 5)
	fmt.Println("after deleted: ", len(month))
	fmt.Println()

	// check whether a key is existed in a map
	// map has 2 return values, one is optional (bool)
	value, isExist := month[5]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item doesn't exist")
	}
	fmt.Println()

	// shorter version
	if value, isExist := month[5]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item doesn't exist")
	}
	fmt.Println()

	// map and slice combined:
	var students = []map[string]string{
		{"ID": "S001", "Name": "Rafael", "Address": "Jakarta"},
		{"ID": "S002", "Name": "Sitta", "Address": "India"},
		{"ID": "S003", "Name": "Jaipur", "Address": "Pakistan"},
	}

	for _, student := range students {
		fmt.Println(student["Name"])
	}
}
