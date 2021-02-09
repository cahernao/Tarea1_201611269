package Handlers

import (
	"net/http"
)

func GetTiendaEspecializadaHandler(w http.ResponseWriter, r *http.Request) {
	var tiendas []EDD.TiendaEspecifica
	for _, v := range TiendasGeneral {
		tiendas = append(tiendas, v)
	}

}
