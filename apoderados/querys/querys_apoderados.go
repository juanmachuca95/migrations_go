package apoderados

var GetApoderados = func() string {
	return "SELECT apoderados.id, apoderados.idPersona, apoderados.created_at, apoderados.updated_at, sas.id AS sas_id FROM `sasapoderados` AS apoderados INNER JOIN saspersonas AS personas ON apoderados.idPersona = personas.id INNER JOIN sasform AS sas ON sas.id = personas.idSAS;"
}

var CreateApoderadosSAS = func() string {
	return "INSERT INTO sas_apoderados (id, sass_id, personas_id, activo, created_at, updated_at) VALUES (?,?,?,?,?,?)"
}
