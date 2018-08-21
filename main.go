package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("")))

	fmt.Println("Hosting on http://localhost:8080")
	fmt.Println("Stop by pressing ^C")
	fmt.Println(http.ListenAndServe(":8080", r))
}
