package users

import (
	"database/sql"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
)

type UsersGateway interface {
	GetUsers() string
}

type UsersService struct {
	*sql.DB
}

func NewUsersGateway() UsersGateway {
	return &UsersService{
		internal.MySQLConnection(),
	}
}

func (s *UsersService) GetUsers() string {

	return "hola"
}
