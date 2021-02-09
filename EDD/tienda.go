package EDD

type Tienda struct {
	nombre, descripcion, contacto string
	calificacion                  int
}

type BuscarTienda struct {
	Departamento string `json:"Departamento"`
	Nombre       string `json:"Nombre"`
	Calificacion int    `json:"Calificacion"`
}

var TiendasGeneral = make(map[string]BuscarTienda)
var NumeroTiendas int
