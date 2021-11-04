package sas

import (
	"database/sql"
	"log"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"

	models "github.com/juanmachuca95/migrations_go/sas/models"
	querys "github.com/juanmachuca95/migrations_go/sas/querys"
)

type SasGateway interface {
	GetSas() ([]models.Sasform, error)
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

func (s *ServiceSas) CreateSAS(sasforms []models.Sasform) (bool, error) {

	return true, nil
}
