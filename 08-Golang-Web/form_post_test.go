package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// firstName := r.PostForm.Get("first_name")
	// lastName := r.PostForm.Get("last_name")

	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")

	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Go&last_name=Lang")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/hello", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
