package sas

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
	"github.com/juanmachuca95/migrations_go/utils"

	models "github.com/juanmachuca95/migrations_go/sas/models"
	querys "github.com/juanmachuca95/migrations_go/sas/querys"
)

type SasGateway interface {
	GetSas() ([]models.Sasform, error)
	CreateSAS([]models.Sasform) (bool, error)

	/*Log register*/
	Logg(string) bool
}

type ServiceSas struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewSasGateway() SasGateway {
	return &ServiceSas{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *ServiceSas) GetSas() ([]models.Sasform, error) {
	stmt, err := s.db.Prepare(querys.GetSas())
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var sasforms []models.Sasform
	var sasform models.Sasform
	for rows.Next() {

		err := rows.Scan(
			&sasform.Id,
			&sasform.Id_Usuario,
			&sasform.Id_Estado,
			&sasform.Razon_Social,
			&sasform.Provincia,
			&sasform.Localidad,
			&sasform.Calle,
			&sasform.Altura,
			&sasform.Piso,
			&sasform.Depto,
			&sasform.Num_Sede,
			&sasform.Duracion,
			&sasform.Peso_Num,
			&sasform.Porcintegr,
			&sasform.Fecha_Presentacion,
			&sasform.Fecha_Aprobacion,
			&sasform.Block,
			&sasform.Created_At,
			&sasform.Updated_At,
			&sasform.Monto_Minimo,
			&sasform.Estatuto_Modelo,
			&sasform.Fecha_Firma,
			&sasform.Fecha_Cierre_Fiscal,
			&sasform.Lugar_Firma,
			&sasform.Tipo_Instrumento,
			&sasform.Fecha_Baja,
			&sasform.Fecha_Rechazo,
			&sasform.Fecha_A_Presentar,
			&sasform.Entidad_Certif,
			&sasform.Bloqueado,
			&sasform.Actividad,
			&sasform.Email,
			&sasform.Telefono,
			&sasform.Completo,
			&sasform.Fecha_Rentas,
			&sasform.Fecha_Boletin,
			&sasform.Nro_Expediente,
			&sasform.Cbu,
			&sasform.Nro_Cta_Bco,
			&sasform.Firmado,
			&sasform.Actividad_Afip,
			&sasform.Motivo_Rechazo,
			&sasform.Entidad_Rechazante,
			&sasform.Pago_Aranceles,
			&sasform.Pago_Integracion_Capitales,
			&sasform.Doc_Escribano_Subido,
			&sasform.Fecha_Ultimo_Pago,
			&sasform.Digitalizado,
			&sasform.Cuit,
			&sasform.Impuesto_Nac,
			&sasform.Sucursal_Banco,
			&sasform.Solicitar_Banco,
			&sasform.Archivo_Rentas,
			&sasform.Expediente_Ciudadano,
			&sasform.Convenio_Rentas,
			&sasform.Comprobante_Integracion,
			&sasform.Puede_Publicar,
			&sasform.Fecha_Retiros_Doc,
			&sasform.Envio_Mail_Sas_fin,
			&sasform.Papelera,
			&sasform.Objeto_Societario,
			&sasform.Baja_En_Afip,
			&sasform.Baja_En_Rentas,
			&sasform.Ultimo_Balance,
			&sasform.Baja_Status,
			&sasform.Baja_Msg)

		if err != nil {
			log.Fatal(err)
		}

		sasforms = append(sasforms, sasform)
	}

	return sasforms, nil
}

/*
	Datos para insertar en sas final
	(id, sas, homonimia, fecha_otorgamiento_instrumento, entidad_certificantes_id, tipo_instrumentos_id, objeto_societarios_id, firma_ciudads_id, fecha_cierre_ejercicios_id,
	capital_minimo, capital_invertido, sede_ciudads_id, calle, altura, piso, dpto, codigo_actividad_economica, regimen_tributario_rentas_id, actividad_rentas, solicita_bc,
	bancos_sucursals_id, created_at, updated_at, descripcion_objeto, estado, comprobante_integracion, users_id, activo, fecha_presentacion, fecha_aprobacion, pagoAranceles)
*/
func (s *ServiceSas) CreateSAS(sasforms []models.Sasform) (bool, error) {

	var objeto_especifico_id = s.GetIdOjetoSocietario("Objeto Específico")
	var objeto_amplio_id = s.GetIdOjetoSocietario("Objeto Amplio")
	var entidad_certificantes_id = 1 // Escribano
	var tipo_instrumentos_id = 1     // Instrumento Privado
	var objeto_societarios_id int
	var firma_ciudads_id int
	var fecha_cierre_ejercicios_id int

	stmt1, err := s.db2.Prepare(querys.CreateSAS())
	if err != nil {
		log.Fatalf("Error al preparar la consulta - error: %s", err)
	}
	defer stmt1.Close()

	for _, value := range sasforms {
		if value.Id_Estado > 2 {

			// Objeto societario
			if value.Objeto_Societario != nil {
				objeto_societarios_id = objeto_especifico_id
			} else {
				objeto_societarios_id = objeto_amplio_id
			}

			// Firma ciudad
			firma_ciudads_id = s.GetFirmaCiudadId(*value.Lugar_Firma)

			// Fecha Cierre Fiscal
			fecha_cierre_ejercicios_id = s.GetFechaCierreFiscal(*value.Fecha_Cierre_Fiscal)

			_, err := stmt1.Exec(value.Id, value.Razon_Social, value.Fecha_Firma, entidad_certificantes_id, tipo_instrumentos_id, objeto_societarios_id, firma_ciudads_id, fecha_cierre_ejercicios_id, value.Monto_Minimo)
			if err != nil {
				errLog := fmt.Sprintf("Ha ocurrido un error al insertar sas id: %d - error: %s", value.Id, err)
				s.Logg(errLog)
			}
		}
	}

	return true, nil
}

func (s *ServiceSas) GetIdOjetoSocietario(objeto string) int {
	var objeto_societario_id int

	stmt2, err := s.db2.Prepare(querys.GetObjetoSocietarioId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta - error: %s", err)
	}

	err = stmt2.QueryRow(objeto).Scan(&objeto_societario_id)
	if err != nil {
		log.Fatalf("Error al obtener id objeto societario: %s - error: %s", objeto, err)
	}

	return objeto_societario_id
}

func (s *ServiceSas) GetFirmaCiudadId(ciudad string) int {
	var ciudads_id int
	stmt, err := s.db2.Prepare(querys.GetFirmaCiudadId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta - error: %s", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(ciudad).Scan(&ciudads_id)
	if err != nil {
		log.Fatalf("Error al obtener id ciudad: %s - error: %s", ciudad, err)
	}

	return ciudads_id
}

func (s *ServiceSas) GetFechaCierreFiscal(fecha string) int {
	var fecha_cierre_ejercicios_id int
	stmt, err := s.db2.Prepare(querys.GetFechaCierreFiscalId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta - error: %s", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(fecha).Scan(&fecha_cierre_ejercicios_id)
	if err != nil {
		s.CreateFechaCierreFiscal(fecha)
		s.GetFechaCierreFiscal(fecha)
	}

	return fecha_cierre_ejercicios_id
}

func (s *ServiceSas) CreateFechaCierreFiscal(fecha string) bool {
	var mes int
	var dia int

	words := strings.Fields(fecha)
	dia, _ = strconv.Atoi(words[0])
	mes = utils.Month(words[2])

	stmt, err := s.db2.Prepare(querys.CreateFechaCierreFiscal())
	if err != nil {
		log.Fatalf("Error al preparar consulta cierre ejercicio fiscal - error: %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(fecha, dia, mes, 1)
	if err != nil {
		log.Fatalf("Error al crear fecha cierre ejercicio: %s - error: %s", fecha, err)
	}

	return true
}

/* Registar los logs error de migrations */
func (s *ServiceSas) Logg(message string) bool {
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
