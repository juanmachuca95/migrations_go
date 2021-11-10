package admins

import "database/sql"

type Admin struct {
	Id               int            `json:"id"`
	Persona_Id       int            `json:"personas_id"`
	Titular          bool           `json:"titular"`
	Inciso           sql.NullString `json:"inciso"`
	Created_At       sql.NullString `json:"created_at"`
	Updated_At       sql.NullString `json:"updated_at"`
	Expuesto_Poli    bool           `json:"expuesto_politicamente"`
	Domicilio_Fiscal sql.NullString `json:"domicilio_fiscal"`
	Admin_Relaciones bool           `json:"admin_relaciones"`
	TelefonoFiscal   sql.NullString `json:"telefono_fiscal"`
	AreaTelefono     sql.NullString `json:"area_telefono"`

	Sas_Id int `json:"sas_id"`
}
