package main

import "fmt"

func main() {
	name := "Cloud"

	switch name {
	case "Cloud":
		fmt.Println("Hello Cloud")
	case "Tifa":
		fmt.Println("Hello Tifa")
	default:
		fmt.Println("Anda siapa?")
	}

	// switch length := len(name); length > 4 {
	// case true:
	// 	fmt.Println("Are you Cloud?")
	// case false:
	// 	fmt.Println("Are you Tifa?")
	// }

	length := len(name)

	switch {
	case length > 4:
		fmt.Println("Are you Cloud?")
	case length > 3:
		fmt.Println("Are you Tifa?")
	default:
		fmt.Println("Who are you?")
	}
}
