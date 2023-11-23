package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Belajar")
	data.PushBack("Golang")
	data.PushBack("Dasar")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	// data.PushFront("Baru")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)
	// fmt.Println(data.Front().Next().Next().Next().Next())

	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
	for el := data.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Value)
	}
}
