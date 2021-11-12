package archivos

import (
	"io"
	"net/http"
	"os"
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
