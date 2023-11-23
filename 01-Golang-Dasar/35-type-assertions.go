package main

import "fmt"

func random() interface{} {
	return "ok"
}

func main() {
	var result interface{} = random()

	// insecure way
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// safer way
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")

	}
}
