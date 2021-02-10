package EDD

type Tienda struct {
	Nombre, Descripcion, Contacto string
	Calificacion                  int
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

func IniciarDepartamento() *Departamento{
	return &Departamento{}
}

var TiendasGeneral = make(map[string]BuscarTienda)
