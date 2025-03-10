package learn_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddelware struct {
	Handler http.Handler
}

func (middleware *LogMiddelware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("after")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)

		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Executed")
		fmt.Fprint(w, "Hellow")
	})
	mux.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("go Executed")
		fmt.Fprint(w, "Hellow go")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panic Executed")
		panic("upss")
	})
	LogMiddelware := &LogMiddelware{
		Handler: mux,
	}
	ErrorHandler := &ErrorHandler{
		Handler: LogMiddelware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
