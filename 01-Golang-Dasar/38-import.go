package main

import (
	"Golang/src/belajar-golang/01-Golang-Dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Cloud")
	// helper.sayGoodBye("Cloud") // error

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

}
