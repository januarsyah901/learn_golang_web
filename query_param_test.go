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

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello, world!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}
func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Janu", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Fprintln(os.Stdout, bodyString)
}
func MultipleQueryParam(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello, %s %s!", firstName, lastName)
}
func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?first_name=Janu&last_name=Akbar", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Fprintln(os.Stdout, bodyString)
}
func MultipleParamaterValue(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}
func TestMultipleParamaterValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Januarsyah&name=Akbar", nil)
	recorder := httptest.NewRecorder()

	MultipleParamaterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Fprintln(os.Stdout, bodyString)
}
