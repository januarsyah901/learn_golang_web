package learn_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText) //manual
	//if err != nil {
	//	panic(err)

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hello Html Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
func TemplateParseFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "hallo guesss")
}
func TestSimpleHTMLTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateParseFile(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "hallo gessss")
}
func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "hallo gessss")
}
func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
