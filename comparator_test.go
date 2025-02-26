package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func ComparatorTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Value": 88,
	})
}

func TestComparatorTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ComparatorTemplate(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
