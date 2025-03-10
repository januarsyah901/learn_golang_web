package learn_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.gohtml", nil)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
	myTemplates.ExecuteTemplate(w, "upload.gohtml", nil)
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/janu.jpeg
var uploadFileTest []byte

func TestUpload(t *testing.T) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	w.WriteField("name", "Januarsyah Akbar")

	file, _ := w.CreateFormFile("file", "gambar.png")
	file.Write(uploadFileTest)
	w.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", w.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)
	bodyRespone, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyRespone))
}
