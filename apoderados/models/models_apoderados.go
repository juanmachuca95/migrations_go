package apoderados

type Apoderado struct {
	Id         int    `json:"id"`
	Persona_Id int    `json:"persona_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`

	Sas_Id int `json:"sas_id"`
}
