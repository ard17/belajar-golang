package main

import "fmt"

func main() {
	var number32 int32 = 131
	var number16 int16 = int16(number32)
	var number8 int8 = int8(number32)

	fmt.Println(number32)
	fmt.Println(number16)
	fmt.Println(number8)

	var name = "Sephiroth"
	var h = name[3]
	var hString string = string(h)

	fmt.Println(name)
	fmt.Println(h)
	fmt.Println(hString)
}
