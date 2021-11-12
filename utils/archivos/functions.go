package archivos

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/minio/minio-go"
)

func DescargarArchivoDeInternet(url string) (string, error) {
	/*
		En este caso voy a descargar una imagen PNG
	*/
	//nombreArchivoSalida := "imagen.png"
	nombreArchivoSalida := "imagen.jpg"

	respuesta, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer respuesta.Body.Close()
	archivoSalida, err := os.Create(nombreArchivoSalida)
	if err != nil {
		return "", err
	}
	defer archivoSalida.Close()
	_, err = io.Copy(archivoSalida, respuesta.Body)
	return nombreArchivoSalida, err
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
