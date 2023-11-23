package main

import "fmt"

func main() {
	var str1 = "Belajar"
	var str2 = "Mengajar"

	var result = str1 == str2
	fmt.Println(result)

	var val1 = 90
	var val2 = 110
	var val3 = 110

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)
	fmt.Println(val2 >= val3)
	fmt.Println(val2 > val3)

}
