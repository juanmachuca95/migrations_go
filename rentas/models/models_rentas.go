package rentas

import "database/sql"

type Renta struct {
	Id         int            `json:"id"`
	Sas_Id     int            `json:"sas_id"`
	Nombre     string         `json:"nombre"`
	File       sql.NullString `json:"file"`
	Ruta       string         `json:"ruta"`
	Mime       string         `json:"mime"`
	Created_At string         `json:"created_at"`
	Updated_At string         `json:"updated_at"`
}
