package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Cloud")
	fmt.Println(result)
	fmt.Println(getGoodBye("Cloud lagi"))
}
