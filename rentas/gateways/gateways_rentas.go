package rentas

import (
	"database/sql"
	"log"

	models "github.com/juanmachuca95/migrations_go/rentas/models"
	querys "github.com/juanmachuca95/migrations_go/rentas/querys"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
)

type RentasGateway interface {
	GetRentas() ([]models.Renta, error)
	CreateRentasSAS([]models.Renta) (bool, error)
}

type ServiceRentas struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewRentasGateway() RentasGateway {
	return &ServiceRentas{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceRentas) GetRentas() ([]models.Renta, error) {
	stmt, err := s.db.Prepare(querys.GetRentas())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get rentas - error: %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta get rentas - error: %s", err)
	}

	var rentas []models.Renta
	var renta models.Renta
	for rows.Next() {
		err = rows.Scan(&renta.Id, &renta.Sas_Id, &renta.Nombre, &renta.File, &renta.Ruta, &renta.Mime, &renta.Created_At, &renta.Updated_At)
		if err != nil {
			log.Fatalf("Error al scanear rows get rentas - error: %s", err)
		}

		rentas = append(rentas, renta)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error en rows get rentas - error: %s", err)
	}

	return rentas, nil
}

func (s *ServiceRentas) CreateRentasSAS(rentas []models.Renta) (bool, error) {

	return true, nil
}
