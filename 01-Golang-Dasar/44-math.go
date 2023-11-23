package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(4.40))
	fmt.Println(math.Floor(4.60))
	fmt.Println(math.Ceil(4.40))
	fmt.Println(math.Max(4.40, 3.6))
	fmt.Println(math.Min(4.40, 30))
}
