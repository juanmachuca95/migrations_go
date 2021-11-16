package archivos

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/minio/minio-go"
)

func DownloadFileOnline(url string, nameFile string) (string, error) {
	nombreArchivoSalida := nameFile + ".jpg"
	log.Printf("La url es: %s", url)
	log.Printf("nombre del archivo: %s", nombreArchivoSalida)
	respuesta, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode != 200 {
		err := errors.New("el archivo no existe")
		return "", err
	}

	data, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	ioutil.WriteFile(nombreArchivoSalida, data, 0666)
	return nombreArchivoSalida, nil
}

func CheckBucket(minio minio.Client, bucketName string) (bool, error) {
	err := minio.MakeBucket(bucketName, os.Getenv("MINIO_LOCATION"))
	if err != nil {
		exists, errBucketExists := minio.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
			return true, nil
		} else {
			return false, err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
		return true, nil
	}
}
