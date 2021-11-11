package apoderados

import (
	"database/sql"
	"log"

	models "github.com/juanmachuca95/migrations_go/apoderados/models"
	querys "github.com/juanmachuca95/migrations_go/apoderados/querys"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
)

type ApoderadoGateway interface {
	GetApoderados() ([]models.Apoderado, error)
	CreateApoderadosSAS([]models.Apoderado) (bool, error)
}

type ServiceApoderado struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewApoderadoGateway() ApoderadoGateway {
	return &ServiceApoderado{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceApoderado) GetApoderados() ([]models.Apoderado, error) {
	stmt, err := s.db.Prepare(querys.GetApoderados())
	if err != nil {
		log.Fatalf("Error al preparar consulta get apoderados - error: %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error al ejecutar consulta get apoderados - error: %s", err)
	}

	var apoderados []models.Apoderado
	var apoderado models.Apoderado
	for rows.Next() {
		err = rows.Scan(&apoderado.Id, &apoderado.Persona_Id, &apoderado.Created_At, &apoderado.Updated_At, &apoderado.Sas_Id)
		if err != nil {
			log.Fatalf("Error al scanear registros rows apoderados - error: %s", err)
		}
		apoderados = append(apoderados, apoderado)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error in rows - error: %s", err)
	}

	return apoderados, nil
}

func (s *ServiceApoderado) CreateApoderadosSAS(apoderados []models.Apoderado) (bool, error) {
	stmt, err := s.db2.Prepare(querys.CreateApoderadosSAS())
	if err != nil {
		log.Fatalf("Error al preparar consulta create apoderados - error: %s", err)
	}
	defer stmt.Close()

	var activo = 1
	for _, value := range apoderados {
		_, err = stmt.Exec(value.Id, value.Sas_Id, value.Persona_Id, activo, value.Created_At, value.Updated_At)
		if err != nil {
			log.Fatalf("Error al ejecutar consulta create apoderados - error: %s", err)
		}
	}

	return true, nil
}
