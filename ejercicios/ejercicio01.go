package ejercicios

import "strconv"

func ConviertoaNumero(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	var valor string
	if err != nil {
		valor = "Hubo un error " + err.Error()
	} else if numero < 100 {
		valor = " menor que 100"
	} else {
		valor = " mayor que 100"
	}
	return numero, valor

}
