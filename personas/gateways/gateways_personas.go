package personas

import (
	"database/sql"
	"log"

	models "github.com/juanmachuca95/migrations_go/personas/models"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
	querys "github.com/juanmachuca95/migrations_go/personas/querys"
)

type PersonasGateway interface {
	GetPersonas() (string, error)
	CreatePersonasSAS([]models.Persona) (bool, error)
	GetDatosPersonaFisica(int) (models.PersonaFisica, error)
}

type ServicePersona struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewPersonaGateway() PersonasGateway {
	return &ServicePersona{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

/*
	Recupera los registros de persona de la base datos justicia
	(id, idTipoDoc, documento, provincia, localidad, calle, altura, piso, depto, block, created_at, updated_at, tipoCuitCuilCdi, CuitCuilCdi, email, telefono);
*/
func (s *ServicePersona) GetPersonas() (string, error) {
	stmt, err := s.db.Prepare(querys.GetPersonas())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al preparar la consulta: %v", err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Ha ocurrido un error al ejecutar la consulta: %v", err)
	}

	var personas []models.Persona
	var persona models.Persona
	for rows.Next() {
		err := rows.Scan(&persona.Id, &persona.Tipo_Doc, &persona.Documento, &persona.Provincia, &persona.Localidad, &persona.Calle, &persona.Altura, &persona.Piso, &persona.Depto, &persona.Block, &persona.Created_At, &persona.Updated_At, &persona.Tipo_Cuit_Cuil_Cdi, &persona.Cuit_Cuil_Cdi, &persona.Email, &persona.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		personas = append(personas, persona)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Ha ocurrido un error al recorrer los datos: %v", err)
	}

	/* Corregir */
	s.CreatePersonasSAS(personas)

	return "hola", nil
}

/*
	Insertar personas de la tabla justicias a (personas)
	Para crear una persona se necesita tener los datos de las tablas (saspersonas, saspersonasjuridicas)

	Datos para crear persona:  (id, cuit, nombre, apellido, documento_tipo, documento_numero, fecha_nacimiento, telefono, email, profesion, paiss_id, estado_civil, ciudads_id, calle, altura, piso, dpto, users_cargo, created_at, updated_at)
*/
func (s *ServicePersona) CreatePersonasSAS(personas []models.Persona) (bool, error) {
	if len(personas) == 0 {
		log.Fatal("No hay personas para registrar")
	}

	stmt, err := s.db2.Prepare(querys.CreatePersonasSAS())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al preparar la consulta: %v", err)
	}
	defer stmt.Close()

	for _, value := range personas {
		var personaFisica models.PersonaFisica

		// Datos persona fisica
		personaFisica, err := s.GetDatosPersonaFisica(value.Id)
		if err != nil {
			log.Fatalf("Ha ocurrido un error al obtener persona fisica id: %d - error: %v", value.Id, err)
		}

		// Create Persona
		_, err = stmt.Exec(value.Id, value.Cuit_Cuil_Cdi, personaFisica.Nombre, personaFisica.Apellido, value.Tipo_Doc, value.Documento, personaFisica.Fecha_Nac, value.Telefono, value.Email, personaFisica.Profesion, 1, personaFisica.Estado_Civil, 1, value.Calle, value.Altura, value.Piso, value.Depto, 1, value.Created_At, value.Updated_At)
		if err != nil {
			log.Fatalf("Ha ocurrido un error al crear persona id: %d - error: %v", value.Id, err)
		}
	}

	return true, nil
}

/* Obtiene la saspersonafisica (estadoCivil, profesion, nombre, apellido, nacionalidad, fechaNac, block, created_at, updated_at) */
func (s *ServicePersona) GetDatosPersonaFisica(personas_id int) (models.PersonaFisica, error) {
	stmt, err := s.db.Prepare(querys.GetDatosPersonaFisica())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al preparar la consulta en tabla saspersonasfisicas: %v", err)
	}

	defer stmt.Close()
	var personaFisica models.PersonaFisica
	err = stmt.QueryRow(personas_id).Scan(&personaFisica.Estado_Civil, &personaFisica.Profesion, &personaFisica.Nombre, &personaFisica.Apellido, &personaFisica.Nacionalidad, &personaFisica.Fecha_Nac, &personaFisica.Block, &personaFisica.Created_At, &personaFisica.Updated_At)
	if err != nil {
		log.Fatalf("Ha ocurrido un error al ejecutar la consulta en tabla saspersonasfisicas: %v", err)
	}

	return personaFisica, nil
}
