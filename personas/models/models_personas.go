package personas

type Persona struct {
	Id                 int     `json:"id"`
	Tipo_Doc           int     `json:"tipo_doc"`
	Tipo_Cuit_Cuil_Cdi int     `json:"tipo_cuit_cuil_cdi"`
	Cuit_Cuil_Cdi      *string `json:"cuit_cuil_cdi"`
	Email              *string `json:"email"`
	Telefono           *string `json:"telefono"`
	Documento          int     `json:"documento"`
	Provincia          *string `json:"provincia"`
	Localidad          *string `json:"localidad"`
	Calle              *string `json:"calle"`
	Altura             *string `json:"altura"`
	Piso               *string `json:"piso"`
	Depto              *string `json:"dpto"`
	Block              int     `json:"block"`
	Created_At         *string `json:"created_at"`
	Updated_At         *string `json:"updated_at"`
}

type PersonaFisica struct {
	Estado_Civil string  `json:"estado_civil"`
	Profesion    string  `json:"profesion"`
	Nombre       string  `json:"nombre"`
	Apellido     string  `json:"apellido"`
	Nacionalidad string  `json:"nacionalidad"`
	Fecha_Nac    string  `json:"fecha_nac"`
	Block        int     `json:"block"`
	Created_At   *string `json:"created_at"`
	Updated_At   *string `json:"updated_at"`
}
