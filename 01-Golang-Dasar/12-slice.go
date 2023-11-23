package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice1 := months[4:7]
	fmt.Println(slice1, "- slice1")
	fmt.Println(len(slice1), "- len slice1")
	fmt.Println(cap(slice1), "- capacity slice1")

	// months[5] = "Diubah"
	// fmt.Println(slice1, "- array diubah")

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months, "- slice diubah")

	slice2 := months[11:]
	fmt.Println(slice2, "- slice2")

	slice3 := append(slice2, "Append")
	fmt.Println(slice3, "- slice3")
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3, "- slice3 diubah")

	fmt.Println(slice2, "- slice2")
	fmt.Println(months, "- months")

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Cloud"
	newSlice[1] = "Strife"

	fmt.Println(newSlice, "- newSlice")
	fmt.Println(len(newSlice), "- len newSlice")
	fmt.Println(cap(newSlice), "- capacity newSlice")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice, "- copySlice")

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniArrayJuga := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArrayJuga)
	fmt.Println(iniSlice)
}
