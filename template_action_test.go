package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func IfActionTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Struct",
		Name:  "Janu",
		Address: Address{
			Street: "JL. Suidrman",
		},
	})
}

func TestIfActionTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	IfActionTemplate(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
