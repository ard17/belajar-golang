package main

import (
	"Golang/src/belajar-golang/01-Golang-Dasar/database"
	"fmt"
	// _ "Golang/src/belajar-golang/01-Golang-Dasar/database" // blank identifier
)

func main() {
	result := database.GetDatabase()

	fmt.Println(result)
}
