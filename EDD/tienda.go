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

type Departamento struct{
	Nombre string `json:",omitempty"`
	Tiendas []Tienda
}

var TiendasGeneral = make(map[string]BuscarTienda)
