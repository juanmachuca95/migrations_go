package mensajes

type infrastructure interface {
	GetGretting() (Mensaje)
}