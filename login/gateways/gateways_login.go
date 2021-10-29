package login

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	internal "github.com/juanmachuca95/migrations_go/internal/database"
	models "github.com/juanmachuca95/migrations_go/login/models"
	querys "github.com/juanmachuca95/migrations_go/login/querys"
	"golang.org/x/crypto/bcrypt"
)

type LoginGateways interface {
	Login(username, password string) (string, error)
}

type LoginService struct {
	*sql.DB
}

func NewLoginGateways() LoginGateways {
	return &LoginService{
		internal.MySQLConnection(),
	}
}

func (s *LoginService) Login(username, password string) (string, error) {
	var user models.Login
	stmt, err := s.Prepare(querys.GetUser())

	if err != nil {
		return "", err
	}

	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&user.Username, &user.Password)
	if err != nil {
		return "", errors.New("Usuario no encontrado.")
	}

	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("Credenciales erroneas.")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24) // A day
	jwtSecret := os.Getenv("TOKEN_KEY")

	token_string, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token_string, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
