package helper

import "fmt"

var version = 1 // tidak bisa diakses dari luar
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// tidak bisa diakses dari luar
func sayGoodBye(name string) {
	fmt.Println("Good bye", name)
}
