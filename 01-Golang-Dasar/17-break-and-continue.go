package main

import "fmt"

func main() {

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Ganjil ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Genap ke", i)
	}

}
