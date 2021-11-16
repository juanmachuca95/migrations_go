package utils

import "log"

func LocalidadStored(localidad string) string {
	var result string
	var localidades = map[string]string{
		"9 De Julio":                  "9 De Julio",
		"Acuña":                       "ACUÑA",
		"Alvear":                      "ALVEAR",
		"Bella Vista":                 "BELLA VISTA",
		"Beron De Astrada":            "BERON DE ASTRADA",
		"Cazadores Correntinos":       "CAZADORES CORRENTINOS",
		"Chavarria":                   "CHAVARRIA",
		"Colonia Carlos Pellegrini":   "COLONIA CARLOS PELLEGRINI",
		"Colonia Libertad":            "COLONIA LIBERTAD",
		"Colonia Liebig's":            "COLONIA LIEBIG'S",
		"Colonia Pando":               "COLONIA PANDO",
		"Concepcion":                  "CONCEPCION",
		"Corrientes":                  "CORRIENTES",
		"Cruz De Los Milagros":        "CRUZ DE LOS MILAGROS",
		"Curuzu Cuatia":               "CURUZU CUATIA",
		"El Sombrero":                 "EL SOMBRERO",
		"Empedrado":                   "EMPEDRADO",
		"Esquina":                     "ESQUINA",
		"Estacion Libertad":           "ESTACION LIBERTAD",
		"Felipe Yofre":                "FELIPE YOFRE",
		"General  Paz":                "GENERAL  PAZ",
		"Gob. Ing. V. Virasoro":       "GOBERNADOR INGENIERO VALENTIN VIRASORO", // Abreviado
		"Gobernador Martinez":         "GOBERNADOR MARTINEZ",
		"Goya":                        "GOYA",
		"Ingeniero Primer Correntino": "INGENIERO PRIMER CORRENTINO",
		"Ita Ibate":                   "ITA IBATE",
		"Itati":                       "ITATI",
		"Ituzaingó":                   "ITUZAINGO",
		"Juan Pujol":                  "JUAN PUJOL",
		"Laguna Brava":                "LAGUNA BRAVA",
		"Las Cuchillas":               "LAS CUCHILLAS",
		"Lavalle":                     "LAVALLE",
		"Libertador":                  "LIBERTADOR",
		"Lomas De Vallejos":           "LOMAS DE VALLEJOS",
		"Loreto":                      "LORETO",
		"Mariano I. Loza":             "MARIANO I. LOZA",
		"Mburucuya":                   "MBURUCUYA",
		"Mercedes":                    "MERCEDES",
		"Mocoreta":                    "MOCORETA",
		"Monte Caseros":               "MONTE CASEROS",
		"Nuestra Señora Del Rosario":  "NUESTRA SEQORA DEL ROSARIO", // Cambios manual
		"Palmar Grande":               "PALMAR GRANDE",
		"Parada Acuña":                "PARADA ACUÑA",
		"Parada Labougle":             "PARADA LABOUGLE",
		"Parada Pucheta":              "PARADA PUCHETA",
		"Paso De La Patria":           "PASO DE LA PATRIA",
		"Paso De Los Libres":          "PASO DE LOS LIBRES",
		"Pedro R. Fernandez":          "PEDRO R. FERNANDEZ",
		"Perugorria":                  "PERUGORRIA",
		"Riachuelo":                   "RIACHUELO",
		"Saladas":                     "SALADAS",
		"San Carlos":                  "SAN CARLOS",
		"San Cayetano":                "SAN CAYETANO",
		"San Cosme":                   "SAN COSME",
		"San Lorenzo":                 "SAN LORENZO",
		"San Luis Del Palmar":         "SAN LUIS DEL PALMAR",
		"San Miguel":                  "SAN MIGUEL",
		"San Roque":                   "SAN ROQUE",
		"Santa Lucia":                 "SANTA LUCIA",
		"Santa Rosa":                  "SANTA ROSA",
		"Santiago Alcorta":            "SANTIAGO ALCORTA",
		"Santo Tome":                  "SANTO TOME",
		"Sauce":                       "SAUCE",
		"Tabay":                       "TABAY",
		"Tacuari":                     "TACUARI",
		"Villa Olivari":               "VILLA OLIVARI",
		"Yahape":                      "YAHAPE",
		"Yapeyú":                      "YAPEYU"}

	for key, value := range localidades {
		if key == localidad {
			result = value
			return result
		}
	}
	return ""
}

func LocalidadesFix(localidad string) string {
	var localidadfixed string
	switch localidad {
	case "Tata Cua":
		localidadfixed = "Tatacuá"
	default:
		localidadfixed = localidad
	}
	return localidadfixed
}

func TipoDocumento(tipoDocId int) string {
	var documento string
	var tipoDocumentos = map[int]string{
		1: "DNI",
		2: "CUIL",
		3: "CUIT",
		4: "LE",
		5: "CDI",
	}

	for index, value := range tipoDocumentos {
		if tipoDocId == index {
			documento = value
		}
	}

	if documento == "" {
		log.Fatal("No se ha podido identificar el documento almacenado")
	}
	return documento
}

func Month(month string) int {
	var number_month int
	var months = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	for index, value := range months {
		if month == value {
			number_month = index
		}
	}

	return number_month
}

func LoteTest() []string {
	var imagenes_test = []string{
		"public/imagenes/usuariosRegistrados/0b68d20f927263cce80acd3e00e22b55_th",
		"public/imagenes/usuariosRegistrados/0b68d20f927263cce80acd3e00e22b55",
		"public/imagenes/usuariosRegistrados/3c51067bcb8c964fe9cdd7d73c6c5fd5_th",
		"public/imagenes/usuariosRegistrados/3c51067bcb8c964fe9cdd7d73c6c5fd5",
		"public/imagenes/usuariosRegistrados/noquieroqueencuentresestaimagen", // failed
	}

	return imagenes_test
}
