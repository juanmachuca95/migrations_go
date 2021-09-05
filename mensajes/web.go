package mensajes


func GetGrettingHandler(w http.ResponseWriter, r *http.Request) Mensaje {

	mensaje := Mensaje {
		mensaje: "Hola mundo desde mi api golang"
	}
	return mensaje
}