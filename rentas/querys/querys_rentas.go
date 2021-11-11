package rentas

var GetRentas = func() string {
	return "SELECT id, idSAS, nombre, file, ruta, mime, created_at, updated_at FROM sasarchivosrentas"
}

var CreateRentasSAS = func() string {
	return "INSERT INTO "
}
