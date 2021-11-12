package personas

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/juanmachuca95/migrations_go/personas/models"
	"github.com/juanmachuca95/migrations_go/utils"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
	querys "github.com/juanmachuca95/migrations_go/personas/querys"
)

type PersonasGateway interface {
	GetPersonas() (bool, error)
	CreatePersonasSAS([]models.Persona) (bool, error)
	GetDatosPersonaFisica(int) (models.PersonaFisica, error)

	/*Log register*/
	Logg(string) bool
}

type ServicePersona struct {
	db  *sql.DB // Justicia
	db2 *sql.DB // Sas
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
func (s *ServicePersona) GetPersonas() (bool, error) {
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

	return true, nil
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
		log.Fatalf("Ha ocurrido un error al preparar la consulta (create persona): %v", err)
	}
	defer stmt.Close()

	stmt2, err := s.db2.Prepare(querys.GetPaisIdPersona())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al preparar la consulta (get pais id persona): %v", err)
	}
	defer stmt.Close()

	stmt3, err := s.db2.Prepare(querys.GetProvinciaIdCiudadIdPersona())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al preparar la consulta (get provincia Id - ciudad Id  de persona) - error : %v", err)
	}
	defer stmt3.Close()

	var tipo_persona = "FISICA"
	for _, value := range personas {
		var personaFisica models.PersonaFisica
		var nacionalidad_id, provincias_id, ciudads_id int

		/*Provincia & ciudad*/
		localidad := utils.LocalidadesFix(*value.Localidad)
		localidad = "%" + localidad + "%"
		err = stmt3.QueryRow(value.Provincia, localidad).Scan(&provincias_id, &ciudads_id)
		if err != nil {
			errLog := fmt.Sprintf("Ha ocurrido un error al intentar obtener provincia %s y ciudad %s de persona id: %d - error: %v", *value.Provincia, *value.Localidad, value.Id, err)
			s.Logg(errLog)
		}

		// Datos persona fisica
		personaFisica, err := s.GetDatosPersonaFisica(value.Id)
		if err != nil {
			errLog := fmt.Sprintf("Ha ocurrido un error al obtener persona fisica - persona id: %d - error: %v", value.Id, err)
			s.Logg(errLog)
		}

		// Nacionalidad
		if personaFisica.Nacionalidad == "Antartida" || personaFisica.Nacionalidad == "Antártida" {
			personaFisica.Nacionalidad = "Argentina"
		}
		err = stmt2.QueryRow(personaFisica.Nacionalidad).Scan(&nacionalidad_id)
		if err != nil {
			errLog := fmt.Sprintf("Ha ocurrido un error al intentar obtener pais %s de persona id: %d - error: %v", personaFisica.Nacionalidad, value.Id, err)
			s.Logg(errLog)
		}

		// Create Persona
		documento_tipo := utils.TipoDocumento(value.Tipo_Doc)
		tipo_cuit_cuil := utils.TipoDocumento(value.Tipo_Cuit_Cuil_Cdi)
		_, err = stmt.Exec(value.Id, value.Cuit_Cuil_Cdi, personaFisica.Nombre, personaFisica.Apellido, documento_tipo, value.Documento, personaFisica.Fecha_Nac, value.Telefono, value.Email, personaFisica.Profesion, nacionalidad_id, personaFisica.Estado_Civil, ciudads_id, value.Calle, value.Altura, value.Piso, value.Depto, 1, value.Created_At, value.Updated_At, tipo_persona, tipo_cuit_cuil)
		if err != nil {
			errLog := fmt.Sprintf("Ha ocurrido un error al crear persona id: %d - error: %v", value.Id, err)
			s.Logg(errLog)
		}

	}

	return true, nil
}

/* Obtiene la saspersonafisica (estadoCivil, profesion, nombre, apellido, nacionalidad, fechaNac, block, created_at, updated_at) */
func (s *ServicePersona) GetDatosPersonaFisica(personas_id int) (models.PersonaFisica, error) {
	stmt, err := s.db.Prepare(querys.GetDatosPersonaFisica())
	if err != nil {
		errLog := fmt.Sprintf("Ha ocurrido un error al preparar la consulta en tabla saspersonasfisicas: %v", err)
		s.Logg(errLog)
	}

	defer stmt.Close()
	var personaFisica models.PersonaFisica
	err = stmt.QueryRow(personas_id).Scan(&personaFisica.Estado_Civil, &personaFisica.Profesion, &personaFisica.Nombre, &personaFisica.Apellido, &personaFisica.Nacionalidad, &personaFisica.Fecha_Nac, &personaFisica.Block, &personaFisica.Created_At, &personaFisica.Updated_At)
	if err != nil {
		errLog := fmt.Sprintf("Ha ocurrido un error al ejecutar la consulta en tabla saspersonasfisicas: %v", err)
		s.Logg(errLog)
	}

	return personaFisica, nil
}

/* Registar los logs error de migrations */
func (s *ServicePersona) Logg(message string) bool {
	stmt, err := s.db2.Prepare("INSERT INTO migrations_logs (error) VALUES (?)")
	if err != nil {
		log.Fatalf("Error al preparar la consulta")
	}
	defer stmt.Close()

	_, err = stmt.Exec(message)
	if err != nil {
		log.Fatalf("Error al registrar el error en la base de datos error: %v", err)
	}
	log.Println(message)

	return true
}
