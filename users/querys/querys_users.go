package users

var GetUsers = func() string {
	return "SELECT id, name, email, password, block, confirmed, confirmation_code, remember_token, created_at, updated_at, apellido, imgUrl, razonSocial, cuit, autorizadoEntrar FROM users"
}

var CreateUsersSAS = func() string {
	return "INSERT INTO users (id, user, name, cuit, email, password, activo, created_at, updated_at, imagen, sidebar) VALUES (?,?,?,?,?,?,?,?,?,?,1)"
}
