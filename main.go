package main

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	tRouter := chi.NewRouter()
	tRouter.Use(middleware.Logger)
	tRouter.Get("/", func(tWriter http.ResponseWriter, tRouter *http.Request) {
		tBody, _ := ioutil.ReadFile("index.html")
		tWriter.Write(tBody)
	})
	http.ListenAndServe(":8080", tRouter)
}
