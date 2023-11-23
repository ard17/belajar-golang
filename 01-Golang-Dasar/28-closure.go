package main

import "fmt"

func main() {
	counter := 1

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)

	name := "Cloud"
	updateName := func() {
		name := "Tifa"
		fmt.Println(name)
		name = "Lockart"
		fmt.Println(name)
	}

	updateName()

	fmt.Println(name)
}
