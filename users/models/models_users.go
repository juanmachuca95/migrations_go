package users

/* Uso un * para poder recibir valores nulos */
type User struct {
	Id                int     `json:"id"`
	Name              *string `json:"name"`
	Apellido          *string `json:"apellido"`
	Email             *string `json:"email"`
	Password          *string `json:"password"`
	Block             int     `json:"block"`
	Confirmed         bool    `json:"confirmed"`
	Confirmation_Code *string `json:"confirmation_cod"`
	Remember_Token    *string `json:"remember_token"`
	Img_Url           *string `json:"img_url"`
	Razon_Social      *string `json:"razon_social"`
	Cuit              *string `json:"cuit"`
	Autorizado_Entrar bool    `json:"autorizado_entrar"`
	Created_At        *string `json:"created_at"`
	Updated_At        *string `json:"updated_at"`
}
