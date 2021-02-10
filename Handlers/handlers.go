package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	. "github.com/chno2016/Tarea1_201611269_30-01-2021/edd"
	"github.com/gorilla/mux"
)

//MiLista una lista parcial
var MiLista []BuscarTienda
var index int

//PostTiendaEspecificaHandler - POST - TiendaEspecifica
func PostTiendaEspecificaHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var t BuscarTienda
	//PODEMOS REGRESAR ESTE VALOR
	//PEND PROBAR REGRESAR
	err := json.Unmarshal(reqBody, &t)
	if err != nil {
		panic(err)
	}

	MiLista = append(MiLista, t)

	json.NewEncoder(w).Encode(t)
	w.WriteHeader(http.StatusCreated)

}

//GetTiendaEspecializadaHandler Funcion para mostrar ese dato con un id
func GetTiendaEspecializadaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k, _ := strconv.Atoi(vars["numero"])
	t := &MiLista[k]
	if t == nil {
		fmt.Printf("No hay contenido que mostrar")
	}
	/*j, err := json.Marshal(t)
	if err != nil {
		//Tener cuidado con esta opcion
		panic(err)
	}*/
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)

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

func GetPrueba1Handler(w http.ResponseWriter, r *http.Request) {

	var miDepa Departamento{}
	var miT []Tienda
	for i := 0; i < 5; i++ {
		t := Tienda{CadenaRandom(), CadenaRandom(), 3}
		miT = append(miT, t)
	}

	miDepa.Nombre = CadenaRandom()
	miDepa.Tiendas=miT

}
func RutaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Carlos Hernandez 20161169")
}

func CadenaRandom() string {
	var output strings.Builder
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	for i := 0; i < 5; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
