package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("./customer-out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Id:       "C030",
		FullName: "Ardiansyah Encoder",
		Address:  "Bekasi",
	}

	encoder.Encode(customer)
}
