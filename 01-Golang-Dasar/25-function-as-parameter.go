package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Asu" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Cloud", spamFilter)
	sayHelloWithFilter("Asu", spamFilter)
}
