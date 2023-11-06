package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Sueldo float32
var Estado bool
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Sueldo = 2750.32
	Estado = true
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Sueldo)
	fmt.Println(Estado)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
