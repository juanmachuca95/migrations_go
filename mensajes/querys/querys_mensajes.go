package mensajes

var GetMensaje = func() string {
	return "SELECT mensaje FROM mensajes WHERE id = ?"
}
