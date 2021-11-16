package socios

import "database/sql"

type Socio struct {
	Id                      int            `json:"id"`
	Persona_Id              int            `json:"persona_id"`
	Sas_id                  int            `json:"sas_id"`
	Cantidad_Acciones_Socio int            `json:"cantidad_acciones_socios"`
	Acciones_Beneficiarios  int            `json:"acciones_beneficiarios"`
	Porcentaje              float64        `json:"porcentaje"`
	Firma                   sql.NullString `json:"firma"`
	Representado            sql.NullInt64  `json:"representado"`
	Created_At              string         `json:"created_at"`
	Updated_At              string         `json:"updated_at"`
}
