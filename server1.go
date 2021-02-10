package main

import (
	"net/http"

	. "github.com/chno2016/Tarea1_201611269_30-01-2021/Handlers"

	"github.com/gorilla/mux"
)

func main() {
	//Nuestro router con strictslach falso que significa que no hay
	//diferencia de un http con o sin slash
	r := mux.NewRouter().StrictSlash(true)

	//endpoint
	//r.HandleFunc("/cargartienda", GetTiendaHandler).Methods("GET")
	//r.HandleFunc("/getArreglo", GetArregloHandler).Methods("GET")
	r.HandleFunc("/TiendaEspecifica", PostTiendaEspecificaHandler).Methods("POST")
	r.HandleFunc("/id/{numero}", GetTiendaEspecializadaHandler).Methods("GET")
	r.HandleFunc("/Prueba1", GetPrueba1Handler).Methods("GET")
	r.HandleFunc("/", RutaInicial)
	//r.HandleFunc("/Eliminar", GetArregloHandler).Methods("GET")
	//r.HandleFunc("/Guardar", PostArregloHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}
