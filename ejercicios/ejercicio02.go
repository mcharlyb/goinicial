package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablasdeMultiplicar() {
	var numero int
	var err error
	scaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un Numero : ")

		if scaner.Scan() {
			numero, err = strconv.Atoi(scaner.Text())
			if err != nil {
				fmt.Println("Ocurrio un error " + err.Error())
			} else {
				for i := 1; i < 11; i++ {
					fmt.Println(numero, " x ", i, " = ", numero*i)
				}
				break
			}
		}
	}
}
