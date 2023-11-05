package main

import (
	"fmt"

	"github.com/mcharlyb/goinicial/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println("el numero es ", texto)

}
