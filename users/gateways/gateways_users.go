package users

import (
	"database/sql"
	"errors"
	"log"
	"os"

	files "github.com/juanmachuca95/migrations_go/utils/archivos"

	internal "github.com/juanmachuca95/migrations_go/internal/database"
	internal2 "github.com/juanmachuca95/migrations_go/internal/database2"
	models "github.com/juanmachuca95/migrations_go/users/models"
	querys "github.com/juanmachuca95/migrations_go/users/querys"
	"github.com/minio/minio-go"
)

type UsersGateway interface {
	GetUsers() ([]models.User, error)
	CreateUsersSAS([]models.User) (bool, error)
}

type UsersService struct {
	db  *sql.DB
	db2 *sql.DB
}

func NewUsersGateway() UsersGateway {
	return &UsersService{
		internal.MySQLConnection(),
		internal2.MySQLConnectionDatabase2(),
	}
}

func (s *UsersService) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := s.db.Query(querys.GetUsers())
	if err != nil {
		log.Fatalf("Ha ocurrido un error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Block, &user.Confirmed, &user.Confirmation_Code, &user.Remember_Token, &user.Created_At, &user.Updated_At, &user.Apellido, &user.Img_Url, &user.Razon_Social, &user.Cuit, &user.Autorizado_Entrar)
		if err != nil {
			log.Fatal("Error al scanear usuario: ", err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

/*
	Inserta los usuarios de la base de datos: Justicia
	a la tabla de destino: sas_golang
	Consulta: InsertUsersSAS

	INSERT INTO sas_golang.users (user, name, cuit, email, password, activo, created_at, updated_at, imagen, sidebar)
*/
func (s *UsersService) CreateUsersSAS(users []models.User) (bool, error) {
	if len(users) == 0 {
		log.Fatal("No usuarios para insertar en esta consulta")
	}

	stmt, err := s.db2.Prepare(querys.CreateUsersSAS())
	if err != nil {
		log.Fatal("Ha ocurrido un error al preparar la consulta")
	}

	defer stmt.Close()
	for _, value := range users {

		// Storage image
		s.StorageImageUser(*value.Img_Url)

		_, err := stmt.Exec(value.Id, value.Apellido, value.Name, value.Cuit, value.Email, value.Password, value.Block, value.Created_At, value.Updated_At, value.Img_Url)
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()
	}

	return true, nil
}

func (s *UsersService) StorageImageUser(image_url string) (bool, error) {

	// Guardar imagen de perfil //https://igpjtesting.corrientes.gob.ar/imagenes/usuariosRegistrados/e1b3724141657679222d55ff4801d748.jpg
	log.Println(image_url)

	log.Fatal("Parar")
	// Minio archivos
	minioClient, err := minio.New(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY_ID"), os.Getenv("MINIO_SECRET_ACCESS_KEY"), true)
	if err != nil {
		log.Fatalf("El minioClient ha arrojado un error: %v", err)
	}

	// Carpeta de almacenamiento en Minio
	bucketName := ""
	_, err = files.CheckBucket(*minioClient, bucketName)
	if err != nil {
		log.Fatalf("Bucket invalido - error: %s", err)
	}

	// Data archivo a almacenar
	objectName := "miprimerimagen.jpg"
	filePath := "imagen.jpg"
	contentType := "image/jpeg"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Fatal("NO EXISTE EL ARCHIVO")
	}

	_, err = minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	return true, nil
}
