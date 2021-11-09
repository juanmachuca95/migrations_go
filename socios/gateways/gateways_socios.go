package socios

import (
	"database/sql"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
)

type SociosGateway interface {
	GetSocios()
}

type ServiceSocios struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewSociosGateway() SociosGateway {
	return &ServiceSocios{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceSocios) GetSocios() {

}
