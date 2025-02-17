package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func ResponeCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		//http.Error(w, "Name is empty", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func TestResponeCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=", nil)
	recorder := httptest.NewRecorder()

	ResponeCode(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Fprintln(os.Stdout, string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
