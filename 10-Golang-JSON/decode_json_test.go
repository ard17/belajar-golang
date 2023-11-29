package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Name":"Ban Aspira","Price":250000,"Images":["https://pkg.go.dev/static/shared/logo/go-blue.svg"],"IsBundling":true}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.Images)
	fmt.Println(product.IsFavorite)
}
