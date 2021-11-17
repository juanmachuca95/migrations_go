package escribanos

var GetEscribanos = func() string {
	return "SELECT id, idSAS, serieFoja, numeroFoja, CUIT, apellido, nombre, puesto, esColegiado, numColegiado, pagoColegio, created_at, updated_at, nroRegistro FROM sasescribanos"
}

var CreateEscribanosSAS = func() string {
	return "INSERT INTO sas_escribanos (id, sas_id, nro_serie, nro_foja, escribanos_id, puesto, esColegiado, numColegiado, pagoColegio, nroRegistro, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
}

var CreateEscribano = func() string {
	return "INSERT INTO escribanos (cuit, apellido, nombre, created_at, updated_at) VALUES (?,?,?,NOW(),NOW())"
}

var GetEscribanoId = func() string {
	return "SELECT id FROM sas_escribanos WHERE cuit=? AND nombre=? AND apellido=?"
}
