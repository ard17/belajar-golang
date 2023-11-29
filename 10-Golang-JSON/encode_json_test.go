package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Name       string
	Price      int
	Images     []string
	IsFavorite bool
}

func TestEncodeJSON(t *testing.T) {
	product := Product{
		Name:       "Ban Aspira",
		Price:      250000,
		Images:     []string{"https://pkg.go.dev/static/shared/logo/go-blue.svg"},
		IsFavorite: true,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
