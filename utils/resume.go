package utils

import (
	"fmt"
	"strconv"
)

func UtilsVar() {
	var x int = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	dato, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(dato)

	/* Mapas */
	mapa := make(map[string]int)
	mapa["Llave"] = 16
	fmt.Println(mapa["Llave"])

}
