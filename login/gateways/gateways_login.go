package login

import (
	"database/sql"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	internal "github.com/juanmachuca95/hexagonal_go/internal/database"
	querys "github.com/juanmachuca95/hexagonal_go/login/querys"
	models "github.com/juanmachuca95/hexagonal_go/users/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginGateways interface {
	Login(username, password string) string
}

type LoginService struct {
	*sql.DB
}

func NewLoginService() LoginGateways {
	return &LoginService{
		internal.MySQLConnection(),
	}
}

func (s *LoginService) Login(username, password string) string {
	var user models.User
	stmt, err := s.Prepare(querys.GetUser())

	if err != nil {
		return err.Error()
	}

	err = stmt.QueryRow(username).Scan(&user.Username, &user.Password, &user.Created_At, &user.Updated_At)
	if err != nil {
		return err.Error()
	}

	if !CheckPasswordHash(password, user.Password) {
		return "Password Incorrect"
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24) // A day
	jwtSecret := os.Getenv("TOKEN_KEY")

	token_string, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return err.Error()
	}

	return token_string
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
