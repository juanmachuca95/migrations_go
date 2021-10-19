package login

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	internal "github.com/juanmachuca95/hexagonal_go/internal/database"
	querys "github.com/juanmachuca95/hexagonal_go/login/querys"
	models "github.com/juanmachuca95/hexagonal_go/users/models"
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
	var user models.User
	stmt, err := s.Prepare(querys.GetUser())

	if err != nil {
		log.Fatalf("Error al preparar la consulta, err: %v\n", err)
	}

	err = stmt.QueryRow(username).Scan(&user.Username, &user.Password, &user.Created_At, &user.Updated_At)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta err: %v", err)
	}

	if !CheckPasswordHash(password, user.Password) {
		log.Fatal("Password incorrecta")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24) // A day
	jwtSecret := os.Getenv("TOKEN_KEY")

	token_string, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal(err)
	}

	return token_string, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
