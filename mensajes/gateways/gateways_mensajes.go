package mensajes

type MensajesGateway interface {
	GetGretting() string
}

type MensajesService struct{}

func NewMensajesGateway() MensajesGateway {
	return &MensajesService{}
}

func (s *MensajesService) GetGretting() string {
	return "hola mundo"
}

/* func main() {
	mensajeService := NewMensajesGateway()
	resp := mensajeService.GetGretting()

	fmt.Println(resp)
}
*/
