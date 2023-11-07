package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablasdeMultiplicar() string {
	var numero int
	var err error
	var texto string
	scaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un Numero : ")

		if scaner.Scan() {
			numero, err = strconv.Atoi(scaner.Text())
			if err != nil {
				fmt.Println("Ocurrio un error " + err.Error())
			} else {
				for i := 1; i < 11; i++ {
					texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
				}

				break
			}
		}
	}

	return texto
}
