package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/data.html"))
	template.ExecuteTemplate(w, "data.html", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Data Map",
	})
}
func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

type Address struct {
	Street string
}
type Data struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/data.html"))
	template.ExecuteTemplate(w, "data.html", Data{
		Title: "Template Data Struct",
		Name:  "Data Struct",
		Address: Address{
			Street: "Jalanan",
		},
	})
}
func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
