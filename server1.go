package main

import (
	"log"
	"net/http"
	"time"

	. "github.com/chno2016/Tarea1_201611269_30-01-2021/Handlers"

	"github.com/gorilla/mux"
)

func main() {
	//Nuestro router con strictslach falso que significa que no hay
	//diferencia de un http con o sin slash
	r := mux.NewRouter().StrictSlash(false)

	//endpoint
	//r.HandleFunc("/cargartienda", GetTiendaHandler).Methods("GET")
	//r.HandleFunc("/getArreglo", GetArregloHandler).Methods("GET")
	r.HandleFunc("/TiendaEspecifica", PostTiendaEspecificaHandler).Methods("POST")
	//r.HandleFunc("/id/{numero}", GetTiendaEspecializadaHandler).Methods("GET")
	//r.HandleFunc("/Eliminar", GetArregloHandler).Methods("GET")
	//r.HandleFunc("/Guardar", PostArregloHandler).Methods("POST")

	server := &http.Server{
		Addr:           "3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando...")
	server.ListenAndServe()
}
