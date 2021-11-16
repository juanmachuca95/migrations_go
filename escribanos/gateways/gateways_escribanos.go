package escribanos

import (
	"database/sql"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	models "github.com/juanmachuca95/migrations_go/models/escribanos"
)

type EscribanoGateway struct {
	GetEscribanos() ([]models.Escribano, error)
}

type ServiceEscribano struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewEscribanoGateway() EscribanoGateway {
	return ServiceEscribano{
		internal.MySQLConnection(),
		internal.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceEscribano) GetEscribanos() ([]models.Escribano, error) {
	var escribanos []models.Escribano

	return escribanos, nil
}
