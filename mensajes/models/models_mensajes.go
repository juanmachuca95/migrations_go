package mensajes

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

func NewMensaje(mensaje string) *Mensaje {
	return &Mensaje{
		Mensaje: mensaje,
	}
}
