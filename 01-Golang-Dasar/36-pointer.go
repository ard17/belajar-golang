package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// pass by value
	// address2 := address1

	// pass by reference / pointer
	address2 := &address1
	// address3 := &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Yogyakarta", "DIY", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address3 := &address1
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	fmt.Println(address4)

	address4.City = "Jayapura"
	fmt.Println(address4)

	alamat := Address{City: "Purwakarta", Province: "Jawa Barat", Country: ""}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}
