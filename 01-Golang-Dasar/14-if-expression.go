package main

import "fmt"

func main() {
	name := "Sephiroth"

	if name == "Cloud" {
		fmt.Println("Hello Cloud")
	} else if name == "Tifa" {
		fmt.Println("Hello Tifa")
	} else {
		fmt.Println("Siapa Nama Anda?")
	}

	if length := len(name); length < 4 {
		fmt.Println("Nama harus lebih dari 3 karakter")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}
