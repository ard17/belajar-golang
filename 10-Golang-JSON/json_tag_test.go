package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

func TestJSONTagEncode(t *testing.T) {
	customer := Customer{
		Id:       "C0001",
		FullName: "Golang JSON",
		Address:  "Bekasi",
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"C0001","full_name":"Golang JSON","address":"Bekasi"}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonByte, customer)

	fmt.Println(customer)
}
