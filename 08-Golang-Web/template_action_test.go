package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/if.html"))
	template.ExecuteTemplate(w, "if.html", Data{
		Title: "Template Action If",
		Name:  "Golang",
	})
}
func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateComparator(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/comparator.html"))
	template.ExecuteTemplate(w, "comparator.html", map[string]interface{}{
		"Title":      "Template Comparator",
		"FinalValue": 23,
	})
}
func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/range.html"))
	template.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Title":   "Template Range",
		"Hobbies": []string{"Gaming", "Coding", "Reading"},
	})
}
func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./templates/with.html"))
	template.ExecuteTemplate(w, "with.html", map[string]interface{}{
		"Title": "Template With",
		"Address": map[string]string{
			"Street": "Jalan Tanah Merah",
			"City":   "Bekasi",
		},
	})
}
func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
