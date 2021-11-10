package admins

import (
	"database/sql"
	"log"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"

	models "github.com/juanmachuca95/migrations_go/admins/models"
	querys "github.com/juanmachuca95/migrations_go/admins/querys"
)

type AdminGateway interface {
	GetAdmins() ([]models.Admin, error)
	CreateAdminsSAS([]models.Admin) (bool, error)
}

type ServiceAdmin struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewAdminGateway() AdminGateway {
	return &ServiceAdmin{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

/**
admin.id, admin.idPersona, admin.titular, admin.inciso, admin.created_at, admin.updated_at,
admin.expuestoPoli, admin.domicilioFiscal, admin.adminRelaciones, admin.telefonoFiscal,
admin.areaTelefono, sas.id AS sas_id
*/
func (s *ServiceAdmin) GetAdmins() ([]models.Admin, error) {
	stmt, err := s.db.Prepare(querys.GetAdmins())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get admins - error: %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta get admins - error: %s", err)
	}

	var admins []models.Admin
	var admin models.Admin
	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Persona_Id, &admin.Titular, &admin.Inciso, &admin.Created_At, &admin.Updated_At,
			&admin.Expuesto_Poli, &admin.Domicilio_Fiscal, &admin.Admin_Relaciones, &admin.TelefonoFiscal,
			&admin.AreaTelefono, &admin.Sas_Id)
		if err != nil {
			log.Fatalf("Error al scanear datos - get admins - error: %s", err)
		}

		admins = append(admins, admin)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error en datos obtenidos get admins - error: %s", err)
	}

	return admins, nil
}

/*
	Datos a insertar :
	id, sass_id, personas_id, administrador, politicamente_expuesto, tipo_politicamente_expuesto,
	activo, created_at, updated_at, administrador_relaciones, domicilio_fiscal
*/
func (s *ServiceAdmin) CreateAdminsSAS(admins []models.Admin) (bool, error) {
	stmt, err := s.db2.Prepare(querys.CreateAdminsSAS())
	if err != nil {
		log.Fatalf("Error al preparar la consulta create admins - error: %s", err)
	}
	defer stmt.Close()

	var activo = 1
	for _, value := range admins {
		_, err = stmt.Exec(value.Id, value.Sas_Id, value.Persona_Id, value.Titular, value.Expuesto_Poli, value.Inciso, activo,
			value.Created_At, value.Updated_At, value.Admin_Relaciones, value.Domicilio_Fiscal)
		if err != nil {
			log.Fatalf("Error al insertar admins id: %v - error: %s", value.Id, err)
		}
	}

	return true, nil
}
