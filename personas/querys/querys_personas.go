package personas

var GetPersonas = func() string {
	return "SELECT id, idTipoDoc, documento, provincia, localidad, calle, altura, piso, depto, block, created_at, updated_at, tipoCuitCuilCdi, CuitCuilCdi, email, telefono FROM saspersonas;"
}

var GetDatosPersonaFisica = func() string {
	return "SELECT estadoCivil, profesion, nombre, apellido, nacionalidad, fechaNac, block, created_at, updated_at FROM saspersonasfisicas WHERE idPersona=?"
}

/*
	Tabla: personas
	cuit, nombre, apellido, documento_tipo, documento_numero, fecha_nacimiento,
	telefono, email, profesion, paiss_id, estado_civil, ciudads_id, calle, altura,
	piso, dpto, users_cargo, created_at, updated_at
*/
var CreatePersonasSAS = func() string {
	return "INSERT INTO personas (id, cuit, nombre, apellido, documento_tipo, documento_numero, fecha_nacimiento, telefono, email, profesion, paiss_id, estado_civil, ciudads_id, calle, altura, piso, dpto, users_cargo, created_at, updated_at) VALUES (?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
}

/* Consultas auxiliares a personas*/
var GetPaisIdPersona = func() string {
	return "SELECT id FROM paiss WHERE pais=? LIMIT 1"
}

var GetProvinciaIdCiudadIdPersona = func() string {
	return "SELECT p.id AS provincias_id, c.id AS ciudads_id FROM provincias AS p INNER JOIN ciudads AS c ON c.provincias_id = p.id WHERE p.provincia=? AND c.ciudad LIKE ?"
}
