package socios

var GetSocios = func() string {
	return "SELECT * FROM `sassocios` INNER JOIN sasbeneficiarios ON sassocios.idPersona = sasbeneficiarios.idPersona;"
}
