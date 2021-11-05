package sas

var GetSas = func() string {
	return "SELECT * FROM sasform"
}

var CreateSAS = func() string {
	return "INSERT INTO sass (id, sas, homonimia, fecha_otorgamiento_instrumento, entidad_certificantes_id, tipo_instrumentos_id, objeto_societarios_id, firma_ciudads_id, fecha_cierre_ejercicios_id, capital_minimo, capital_invertido, sede_ciudads_id, calle, altura, piso, dpto, codigo_actividad_economica, regimen_tributario_rentas_id, actividad_rentas, solicita_bc, bancos_sucursals_id, created_at, updated_at, descripcion_objeto, estado, comprobante_integracion, users_id, activo, fecha_presentacion, fecha_aprobacion, pagoAranceles) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
}

/*
	Consultas auxiliares:
	Obtener id objeto societario - Específico o Amplio
	Obtener id ciudad de firma
*/
var GetObjetoSocietarioId = func() string {
	return "SELECT id FROM objeto_societarios WHERE objeto_societario=?"
}

var GetFirmaCiudadId = func() string {
	return "SELECT id FROM ciudads WHERE ciudad=?"
}

var GetFechaCierreFiscalId = func() string {
	return "SELECT id FROM fecha_cierre_ejecicios WHERE fecha_cierre_ejercicio=?"
}

var CreateFechaCierreFiscal = func() string {
	return "INSERT INTO fecha_cierre_ejercicios (fecha_cierre_ejercicio, dia, mes, activo) VALUES (?,?,?,?)"
}