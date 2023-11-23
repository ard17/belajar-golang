package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var cloud Customer
	cloud.Name = "Cloud"
	cloud.Address = "Ciamis"
	cloud.Age = 25

	fmt.Println(cloud)
	fmt.Println(cloud.Name)
	fmt.Println(cloud.Address)
	fmt.Println(cloud.Age)
	cloud.sayHello()

	tifa := Customer{
		Name:    "Tifa",
		Address: "Bandung",
		Age:     24,
	}
	fmt.Println(tifa)
	tifa.sayHello()

	sephiroth := Customer{"Sephiroth", "Jakarta", 27}
	fmt.Println(sephiroth)

}
