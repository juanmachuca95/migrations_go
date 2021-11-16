package socios

import (
	"database/sql"
	"log"

	models "github.com/juanmachuca95/migrations_go/socios/models"
	querys "github.com/juanmachuca95/migrations_go/socios/querys"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
)

type SociosGateway interface {
	GetSocios() ([]models.Socio, error)
	CreateSociosSAS(socios []models.Socio) (bool, error)
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

func (s *ServiceSocios) GetSocios() ([]models.Socio, error) {
	stmt, err := s.db.Prepare(querys.GetSocios())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get socios - error: %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error al ejecutar al consulta get socios - error: %s", err)
	}

	var socios []models.Socio
	var socio models.Socio
	// socios.idPersona, beneficiarios.idPersona, sas.id, socios.cantAcciones,
	// beneficiarios.acciones, beneficiarios.porcentaje, socios.firma, socios.representado,
	// socios.created_at, socios.updated_at;
	for rows.Next() {
		err = rows.Scan(&socio.Id, &socio.Persona_Id, &socio.Sas_id, &socio.Cantidad_Acciones_Socio, &socio.Acciones_Beneficiarios,
			&socio.Porcentaje, &socio.Firma, &socio.Representado, &socio.Created_At, &socio.Updated_At)
		if err != nil {
			log.Fatalf("Error al scannear get socios - error: %s", err)
		}

		socios = append(socios, socio)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error en rows get socios - error: %s", err)
	}

	return socios, nil
}

func (s *ServiceSocios) CreateSociosSAS(socios []models.Socio) (bool, error) {
	stmt, err := s.db2.Prepare(querys.CreateSociosSAS())
	if err != nil {
		log.Fatalf("Error al preparar la consulta create socios - error: %s", err)
	}
	defer stmt.Close()

	//(id, carga_capital_invertido, porcentaje, created_at, updated_at, sass_id, personas_id)
	for _, value := range socios {
		_, err = stmt.Exec(value.Id, value.Cantidad_Acciones_Socio, value.Porcentaje, value.Created_At, value.Updated_At, value.Sas_id, value.Persona_Id)
		if err != nil {
			log.Fatalf("Error al ejecutar la consulta - error: %s", err)
		}
	}

	return true, nil
}
