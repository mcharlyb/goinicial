package main

import (
	"fmt"

	"github.com/mcharlyb/goinicial/ejercicios"
)

func main() {
	numero, texto := ejercicios.ConviertoaNumero("88")
	fmt.Println(numero)
	fmt.Println("el numero es ", texto)

}
