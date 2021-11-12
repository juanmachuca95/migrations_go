package main

import (
	"errors"
	"log"
	"os"

	"github.com/minio/minio-go"
)

func main() {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalf("El minioClient ha arrojado un error: %v", err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	// Make a new bucket called mymusic.
	bucketName := "sasusersadmin"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "miprimerimagen.jpg"
	filePath := "imagen.jpg"
	contentType := "image/jpeg"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		log.Fatal("NO EXISTE EL ARCHIVO")
	}

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info)

	/* err := godotenv.Load()
	if err != nil {
		log.Fatal("No se han podido cargar las variables de entorno.")
	}

	port := os.Getenv("PORT")
	port2 := os.Getenv("PORT2")

	if port == "" {
		port = "8080"
	}

	if port2 == "" {
		port2 = "8081"
	}

	fmt.Println(port)
	api.Start(port, port2) */

}
