package escribanos

import (
	"database/sql"
	"log"

	models "github.com/juanmachuca95/migrations_go/escribanos/models"
	querys "github.com/juanmachuca95/migrations_go/escribanos/querys"
	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
)

type EscribanoGateway interface {
	GetEscribanos() ([]models.EscribanoSAS, error)
	CreateEscribanoSAS(escribanos []models.EscribanoSAS) (bool, error)
}

type ServiceEscribano struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewEscribanoGateway() EscribanoGateway {
	return &ServiceEscribano{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceEscribano) GetEscribanos() ([]models.EscribanoSAS, error) {
	stmt, err := s.db.Prepare(querys.GetEscribanos())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get escribano - error: %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta get escribanos - error: %s", err)
	}

	var escribanosSas []models.EscribanoSAS
	var escribanoSas models.EscribanoSAS
	var escribano models.Escribano
	for rows.Next() {
		err := rows.Scan(&escribanoSas.Id, &escribanoSas.Sas_Id, &escribanoSas.Serie_Foja, &escribanoSas.Numero_Foja, &escribano.Cuit, &escribano.Apellido, &escribano.Nombre, &escribanoSas.Puesto, &escribanoSas.Es_Colegiado, &escribanoSas.Num_Colegiado, &escribanoSas.Pago_Colegio, &escribanoSas.Created_At, &escribanoSas.Updated_At, &escribanoSas.Nro_Registro)
		if err != nil {
			log.Fatalf("Error al scanner los registros get escribanos - error: %s", err)
		}

		escribanosSas = append(escribanosSas, escribanoSas)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error en rows get escribanos - error %s", err)
	}

	return escribanosSas, nil
}

func (s *ServiceEscribano) CreateEscribanoSAS(escribanos []models.EscribanoSAS) (bool, error) {
	stmt, err := s.db2.Prepare(querys.CreateEscribanosSAS())
	if err != nil {
		log.Fatalf("Error al preparar la consulta create escribano sas - error: %s", err)
	}
	defer stmt.Close()

	for _, value := range escribanos {
		var escribanos_id sql.NullInt64
		if value.Nombre.Valid && value.Apellido.Valid && value.Cuit.Valid {
			_, _ = s.CreateEscribano(value.Escribano)
			escribanos_id = s.GetEscribanoId(value.Escribano)
		} else {
			escribanos_id = sql.NullInt64{}
		}

		// (id, sas_id, nro_serie, nro_foja, escribanos_id, puesto, esColegiado,
		// numColegiado, pagoColegiado, nroRegistro, created_at, updated_at)
		_, err = stmt.Exec(value.Id, value.Sas_Id, value.Serie_Foja, value.Numero_Foja, escribanos_id, value.Puesto,
			value.Es_Colegiado, value.Num_Colegiado, value.Pago_Colegio, value.Nro_Registro, value.Created_At, value.Updated_At)
		if err != nil {
			log.Fatalf("Error al ejecutar la consulta create escribano sas - error: %s", err)
		}
	}

	return true, nil
}

func (s *ServiceEscribano) CreateEscribano(escribano models.Escribano) (bool, error) {
	stmt, err := s.db2.Prepare(querys.CreateEscribano())
	if err != nil {
		log.Fatalf("Error al preparar la consulta create escribano - error: %s", err)
	}
	defer stmt.Close()

	//(cuit, apellido, nombre, created_at, updated_at)
	_, err = stmt.Exec(escribano.Cuit, escribano.Apellido)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta create escribano - error: %s", err)
	}

	return true, nil
}

func (s *ServiceEscribano) GetEscribanoId(escribano models.Escribano) sql.NullInt64 {
	stmt, err := s.db.Prepare(querys.GetEscribanoId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get escribano id - error: %s", err)
	}
	defer stmt.Close()

	var escribanos_id sql.NullInt64
	err = stmt.QueryRow(escribano.Cuit, escribano.Apellido, escribano.Nombre).Scan(&escribanos_id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta get escribano id - error: %s", err)
	}

	return escribanos_id
}
