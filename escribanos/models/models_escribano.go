package escribanos

import "database/sql"

type EscribanoSAS struct {
	Id            int            `json:"id"`
	Sas_Id        int            `json:"sas_id"`
	Serie_Foja    sql.NullString `json:"seria_foja"`
	Numero_Foja   sql.NullString `json:"numero_foja"`
	Puesto        sql.NullString `json:"puesto"`
	Es_Colegiado  int            `json:"es_colegiado"`
	Num_Colegiado sql.NullString `json:"num_colegiado"`
	Pago_Colegio  int            `json:"pago_colegiado"`
	Created_At    string         `json:"created_at"`
	Updated_At    string         `json:"updated_at"`
	Nro_Registro  string         `json:"nro_registro"`

	Escribano
}

type Escribano struct {
	Cuit     sql.NullString `json:"cuit"`
	Apellido sql.NullString `json:"apellido"`
	Nombre   sql.NullString `json:"nombre"`
}
