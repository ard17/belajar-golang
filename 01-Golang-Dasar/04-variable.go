package main

import "fmt"

func main() {
	var name string

	name = "Cloud Strife"
	fmt.Println(name)

	name = "Serphiroth"
	fmt.Println(name)

	var anotherName = "Noctis"
	// anotherName = 13  // Error Data Type
	fmt.Println(anotherName)

	var age = 25
	fmt.Println(age)

	country := "Sleeping Forest"
	fmt.Println(country)

	var (
		firstName = "Cloud"
		lastName  = "Strife"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
