package main

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, values []interface{}) {
	outputFile, outputError := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if outputError != nil {
		fmt.Println("Terjadi kesalahan saat membuka file")
		return
	}
	defer outputFile.Close()

	for _, value := range values {
		fmt.Fprintln(outputFile, value) // print values to f, one per line
	}

	_, err := outputFile.WriteString() // input from slice
	if err != nil {
		return
	}
	err = outputFile.Sync()
	if err != nil {
		return
	}
}
