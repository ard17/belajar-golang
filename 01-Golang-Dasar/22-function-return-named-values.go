package main

import "fmt"

func getFullName() (firstName, lastName string) {

	firstName = "Cloud"
	lastName = "Strife"
	return

}

func main() {
	a, b := getFullName()

	fmt.Println(a)
	fmt.Println(b)

}
