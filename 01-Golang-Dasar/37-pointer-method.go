package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) UpdateName() {
	man.Name = "Mr. " + man.Name
}

func main() {
	cloud := Man{"Cloud"}

	cloud.UpdateName()

	fmt.Println(cloud.Name)
}
