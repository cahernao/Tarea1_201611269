package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func rutaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Carlos Hernandez 201611269")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rutaInicial)
	http.ListenAndServe(":3210", router)
}
