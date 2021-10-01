package mensajes

import (
	"database/sql"
	"log"

	querys "github.com/juanmachuca95/hexagonal_go/mensajes/querys"
)

type MensajesGateway interface {
	GetGretting() string
}

type MensajesService struct {
	sql.DB
}

func NewMensajesGateway() MensajesGateway {
	return &MensajesService{}
}

func (s *MensajesService) GetGretting() string {

	var mensaje string
	err := s.Ping()
	if err != nil {
		log.Fatal("acá estamos")
		log.Fatal(err.Error())
	}

	stmt, err := s.Prepare(querys.GetMensaje())

	if err != nil {
		log.Fatal(err.Error())
	}

	err = stmt.QueryRow(1).Scan(&mensaje)
	if err != nil {
		log.Fatal(err.Error())
	}

	return mensaje
}

/* func main() {
	mensajeService := NewMensajesGateway()
	resp := mensajeService.GetGretting()

	fmt.Println(resp)
}
*/
