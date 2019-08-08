package main

import (
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}
	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	port := ":8080"
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}

	log.Println("Listening on port: ", port)
}

func index(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "index", nil)
}
