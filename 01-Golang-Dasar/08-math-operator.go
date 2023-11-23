package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 9
	var c = a * b
	fmt.Println(c)

	var (
		d = 10
		e = 2
		f = d / e
	)
	fmt.Println(f)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -99
	var positive = +99
	fmt.Println(negative)
	fmt.Println(positive)
}
