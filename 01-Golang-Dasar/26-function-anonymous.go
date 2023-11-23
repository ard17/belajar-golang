package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Sephiroth"
	}

	registerUser("Cloud", blacklist)
	registerUser("Sephiroth", blacklist)

	registerUser("Sephiroth", func(name string) bool {
		return name == "Tifa"
	})
	registerUser("Tifa", func(name string) bool {
		return name == "Tifa"
	})

}
