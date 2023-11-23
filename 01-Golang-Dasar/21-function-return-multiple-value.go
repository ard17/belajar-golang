package main

import "fmt"

func getFullName() (string, string) {
	return "Cloud", "Strife"
}

func main() {
	firstname, lastname := getFullName()

	fmt.Println(firstname, lastname)

	namaDepan, _ := getFullName()
	fmt.Println(namaDepan)

	_, namaBelakang := getFullName()
	fmt.Println(namaBelakang)

}
