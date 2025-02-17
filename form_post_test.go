package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func TestFromPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Januarsyah&last_name=Akbar")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Fprintln(os.Stdout, string(body))
}
