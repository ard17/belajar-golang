package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Cloud"
	names[1] = "Strife"
	names[2] = "Sephiroth"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [4]int{
		90,
		85,
		70,
		95,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println(len(values))

	var array [10]string

	fmt.Println(len(array))
}
