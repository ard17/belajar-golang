package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Belajar Golang", "Go"))
	fmt.Println(strings.Split("Belajar Golang", " "))
	fmt.Println(strings.ToLower("Belajar Golang"))
	fmt.Println(strings.ToUpper("Belajar Golang"))
	fmt.Println(strings.Trim("    Belajar Golang    ", " "))
	fmt.Println(strings.ReplaceAll("Belajar Golang Golang", "Go", "Gu"))
}
