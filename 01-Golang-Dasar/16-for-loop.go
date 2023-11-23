package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Cloud", "Strife", "Sephiroth", "Tifa", "Lockhart"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	for _, value := range slice {
		fmt.Println("name =", value)
	}

	person := map[string]string{
		"name":  "Cloud",
		"title": "Swordsman",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
