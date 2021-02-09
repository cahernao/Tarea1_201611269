package Handlers

import (
	"encoding/json"
	"net/http"

	. "github.com/chno2016/Tarea1_201611269_30-01-2021/edd"
)

//PostTiendaEspecificaHandler - POST - TiendaEspecifica
func PostTiendaEspecificaHandler(w http.ResponseWriter, r *http.Request) {
	var t BuscarTienda
	//PODEMOS REGRESAR ESTE VALOR
	//PEND PROBAR REGRESAR
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err)
	}
	EDD.NumeroTiendas++
}

/*


//ES EL GET PERO ESTA CRUZADO XD
func PostTiendaEspecificaHandler(w http.ResponseWriter, r *http.Request) {
	var tiendas []BuscarTienda
	for _, v := range TiendasGeneral {
		tiendas = append(tiendas, v)
	}

	//Cabezera de respuesta
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(TiendasGeneral)
	if err != nil {
		//Tener cuidado con esta opcion
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

**/
