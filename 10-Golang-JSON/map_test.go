package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Ban Aspira",
		"price": 200000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestMapDecode(t *testing.T) {
	jsonString := `{"name":"Ban Aspira","id":"P0001","price":20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Printf("%.0f\n", result["price"])
}
