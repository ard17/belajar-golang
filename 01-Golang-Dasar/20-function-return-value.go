package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello aja"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Cloud")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
