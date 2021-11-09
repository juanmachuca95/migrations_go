package sas

import (
	"database/sql"
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
		//  id, idUsuario, idEstado, razonSocial, provincia, localidad, calle, altura, piso, depto, numSede, duracion, pesoNum, porcIntegr,
		// fechaPresentacion, fechaAprobacion, block, created_at, updated_at, montoMinimo, estatutoModelo,
		// fechaFirma, fechaCierreFiscal, lugarFirma, tipoInstrumento, fechaBaja, fechaRechazo, fechaAPresentar,
		// entidadCertif, bloqueado, actividad, email, telefono, completo, fechaRentas, fechaBoletin,
		// nroExpediente, cbu, NroCtaBco, firmado, actividadAFIP, motivoRechazo, entidadRechazante,  pagoAranceles,
		// pagoIntegracionCapitales, DocEscribanoSubido, fechaUltimoPago, digitalizado, CUIT, impuestoNac,
		// sucursalBanco, solicitarBanco, archivoRentas, ExpedienteCiudadano, ConvenioRentas, comprobanteIntegracion,
		// PuedePublicar, fechaRetiroDoc, EnvioMailSASfin, papelera, objeto_societario, baja_en_afip,
		// baja_en_rentas, ultimo_balance, baja_status, baja_msg
		err := rows.Scan(
			&sasform.Id, &sasform.Id_Usuario, &sasform.Id_Estado, &sasform.Razon_Social, &sasform.Provincia, &sasform.Localidad,
			&sasform.Calle, &sasform.Altura, &sasform.Piso, &sasform.Depto, &sasform.Num_Sede, &sasform.Duracion, &sasform.Peso_Num, &sasform.Porcintegr,
			&sasform.Fecha_Presentacion, &sasform.Fecha_Aprobacion, &sasform.Block, &sasform.Created_At, &sasform.Updated_At, &sasform.Monto_Minimo, &sasform.Estatuto_Modelo,
			&sasform.Fecha_Firma, &sasform.Fecha_Cierre_Fiscal, &sasform.Lugar_Firma, &sasform.Tipo_Instrumento, &sasform.Fecha_Baja, &sasform.Fecha_Rechazo, &sasform.Fecha_A_Presentar,
			&sasform.Entidad_Certif, &sasform.Bloqueado, &sasform.Actividad, &sasform.Email, &sasform.Telefono, &sasform.Completo, &sasform.Fecha_Rentas, &sasform.Fecha_Boletin,
			&sasform.Nro_Expediente, &sasform.Cbu, &sasform.Nro_Cta_Bco, &sasform.Firmado, &sasform.Actividad_Afip, &sasform.Motivo_Rechazo, &sasform.Entidad_Rechazante, &sasform.Pago_Aranceles,
			&sasform.Pago_Integracion_Capitales, &sasform.Doc_Escribano_Subido, &sasform.Fecha_Ultimo_Pago, &sasform.Digitalizado, &sasform.Cuit, &sasform.Impuesto_Nac,
			&sasform.Sucursal_Banco, &sasform.Solicitar_Banco, &sasform.Archivo_Rentas, &sasform.Expediente_Ciudadano, &sasform.Convenio_Rentas, &sasform.Comprobante_Integracion,
			&sasform.Puede_Publicar, &sasform.Fecha_Retiros_Doc, &sasform.Envio_Mail_Sas_fin, &sasform.Papelera, &sasform.Objeto_Societario, &sasform.Baja_En_Afip,
			&sasform.Baja_En_Rentas, &sasform.Ultimo_Balance, &sasform.Baja_Status, &sasform.Baja_Msg)

		if err != nil {
			log.Fatal(err)
		}

		sasforms = append(sasforms, sasform)
	}

	return sasforms, nil
}

/*
	Datos para insertar en sas final
	(id, sas, homonimia, fecha_otorgamiento_instrumento,
	entidad_certificantes_id, tipo_instrumentos_id, objeto_societarios_id, firma_ciudads_id,
	fecha_cierre_ejercicios_id, capital_minimo, capital_invertido, sede_ciudads_id,
	calle, altura, piso, dpto,
	codigo_actividad_economica, regimen_tributario_rentas_id, actividad_rentas, solicita_bc, bancos_sucursals_id,
	created_at, updated_at, descripcion_objeto, estado,
	comprobante_integracion, users_id, activo, fecha_presentacion,
	fecha_aprobacion, pagoAranceles)
*/
func (s *ServiceSas) CreateSAS(sasforms []models.Sasform) (bool, error) {
	var objeto_especifico_id = s.GetIdOjetoSocietario("Objeto Específico")
	var objeto_amplio_id = s.GetIdOjetoSocietario("Objeto Amplio")
	var entidad_certificantes_id = 1 // Escribano
	var tipo_instrumentos_id = 1     // Instrumento Privado
	var homonimia = 1                // Homonimia true
	var (
		objeto_societarios_id,
		regimen_tributario_rentas_id int
	)

	var (
		firma_ciudads_id,
		fecha_cierre_ejercicios_id,
		bancos_sucursals_id,
		sede_ciudads_id sql.NullInt64
	)

	/*Prepare consultas*/
	stmt1, err := s.db2.Prepare(querys.CreateSAS())
	if err != nil {
		log.Fatalf("Error al preparar la consulta create SAS - error: %s", err)
	}
	defer stmt1.Close()

	stmt_ciudad, err := s.db2.Prepare(querys.GetCiudadId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get ciudad id - error: %s", err)
	}
	defer stmt_ciudad.Close()

	stmt_fecha_cierre, err := s.db2.Prepare(querys.GetFechaCierreFiscalId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get fecha cierre fiscal id - error: %s", err)
	}
	defer stmt_fecha_cierre.Close()

	stmt_fecha_cierre1, err := s.db2.Prepare(querys.CreateFechaCierreFiscal())
	if err != nil {
		log.Fatalf("Error al preparar consulta cierre ejercicio fiscal - error: %s", err)
	}
	defer stmt_fecha_cierre1.Close()

	stmt_regimen, err := s.db2.Prepare(querys.GetRegimenTributarioId())
	if err != nil {
		log.Fatalf("Error al preparar consulta get regimen tributario id - error: %s", err)
	}
	defer stmt_regimen.Close()

	stmt_banco, err := s.db2.Prepare(querys.GetBancoSucursalId())
	if err != nil {
		log.Fatalf("Error al preparar consulta get banco sucursal id - error: %s", err)
	}
	defer stmt_banco.Close()

	for _, value := range sasforms {
		log.Println(value.Id_Estado, value.Id)
		if value.Id_Estado <= 6 {

			// Objeto societario
			if value.Objeto_Societario.String != "" {
				objeto_societarios_id = objeto_especifico_id
			} else {
				objeto_societarios_id = objeto_amplio_id
			}

			// Firma ciudad
			firma_ciudads_id = s.GetCiudadId(value.Lugar_Firma.String, stmt_ciudad)

			// Fecha Cierre Fiscal
			if value.Fecha_Cierre_Fiscal.String != "" {
				fecha_cierre_ejercicios_id, err = s.GetFechaCierreFiscal(value.Fecha_Cierre_Fiscal.String, stmt_fecha_cierre)
				if err != nil {
					s.CreateFechaCierreFiscal(value.Fecha_Cierre_Fiscal.String, stmt_fecha_cierre1)
					fecha_cierre_ejercicios_id, _ = s.GetFechaCierreFiscal(value.Fecha_Cierre_Fiscal.String, stmt_fecha_cierre)
				}
			}

			// Sede Ciudad SAS
			sede_ciudads_id = s.GetCiudadId(value.Localidad.String, stmt_ciudad)

			// Regimen tributario
			regimen_tributario_rentas_id = s.GetRegimenTributarioId(value.Convenio_Rentas.String, stmt_regimen)

			// Sucursales banco
			bancos_sucursals_id = s.GetBancoSucursalId(value.Sucursal_Banco.String, stmt_banco)

			_, err = stmt1.Exec(
				value.Id, value.Razon_Social, homonimia, value.Fecha_Firma,
				entidad_certificantes_id, tipo_instrumentos_id, objeto_societarios_id, firma_ciudads_id,
				fecha_cierre_ejercicios_id, value.Monto_Minimo, value.Peso_Num, sede_ciudads_id,
				value.Calle, value.Altura, value.Piso, value.Depto,
				value.Actividad_Afip, regimen_tributario_rentas_id, value.Actividad, value.Solicitar_Banco,
				bancos_sucursals_id, value.Created_At, value.Updated_At, value.Objeto_Societario,
				value.Id_Estado, value.Comprobante_Integracion, value.Id_Usuario, value.Block,
				value.Fecha_Presentacion, value.Fecha_Aprobacion, value.Pago_Aranceles)

			if err != nil {
				log.Fatalf("Ha ocurrido un error al insertar sas id: %d - error: %s", value.Id, err)
			}
		}
	}

	return true, nil
}

func (s *ServiceSas) GetIdOjetoSocietario(objeto string) int {
	if objeto == "" {
		log.Fatalf("El parametro objeto esta vacio param: %s", objeto)
	}
	var objeto_societario_id int

	stmt2, err := s.db2.Prepare(querys.GetObjetoSocietarioId())
	if err != nil {
		log.Fatalf("Error al preparar la consulta get objeto societario id - error: %s", err)
	}

	err = stmt2.QueryRow(objeto).Scan(&objeto_societario_id)
	if err != nil {
		log.Fatalf("Error al obtener id objeto societario: %s - error: %s", objeto, err)
	}

	return objeto_societario_id
}

func (s *ServiceSas) GetCiudadId(ciudad string, stmt *sql.Stmt) sql.NullInt64 {
	var ciudads_id sql.NullInt64
	if ciudad == "" {
		log.Println("Este registro no contiene ciudad asignada")
		return ciudads_id
	}

	if ciudad == "Tata Cua" {
		ciudad = "Tatacuá"
	}

	err := stmt.QueryRow(ciudad).Scan(&ciudads_id)
	if err != nil {
		log.Printf("Error al obtener id ciudad: %s - error: %s\n", ciudad, err)
	}

	return ciudads_id
}

func (s *ServiceSas) GetFechaCierreFiscal(fecha string, stmt *sql.Stmt) (sql.NullInt64, error) {
	var fecha_cierre_ejercicios_id sql.NullInt64
	err := stmt.QueryRow(fecha).Scan(&fecha_cierre_ejercicios_id)
	if err != nil {
		log.Printf("No hay registros de fecha cierre ejercicio para %v - error: %s", fecha, err)
		return fecha_cierre_ejercicios_id, err
	}

	return fecha_cierre_ejercicios_id, nil
}

func (s *ServiceSas) GetRegimenTributarioId(regimen string, stmt *sql.Stmt) int {
	var options = map[string]string{"1": "Régimen convenio multilateral", "0": "Régimen local"}
	var regimen_tributario_rentas_id int

	err := stmt.QueryRow(options[regimen]).Scan(&regimen_tributario_rentas_id)
	if err != nil {
		log.Fatalf("Error al obtener id regimen tributario: %s - error: %s", regimen, err)
	}

	return regimen_tributario_rentas_id
}

func (s *ServiceSas) GetBancoSucursalId(banco string, stmt *sql.Stmt) sql.NullInt64 {
	var bancos_sucursals_id sql.NullInt64
	if banco == "" {
		log.Println("Esta sas no tiene banco registrado")
		return bancos_sucursals_id
	}

	err := stmt.QueryRow(banco).Scan(&bancos_sucursals_id)
	if err != nil {
		log.Printf("Error al obtener id bancos sucursales: %s - error: %s\n", banco, err)
	}

	return bancos_sucursals_id
}

func (s *ServiceSas) CreateFechaCierreFiscal(fecha string, stmt *sql.Stmt) bool {
	var mes, dia int

	words := strings.Fields(fecha)
	dia, _ = strconv.Atoi(words[0])
	mes = utils.Month(words[2])

	_, err := stmt.Exec(fecha, dia, mes, 1)
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
