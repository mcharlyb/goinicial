package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mcharlyb/goinicial/ejercicios"
)

var fileName string = "./files/txt/tablas.txt"

func GrabaTabla() {
	texto := ejercicios.TablasdeMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ocurrio un error " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.TablasdeMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Ocurrio un error al Abrir el Archivo")
		return
	}
}

func Append(filen string, tex string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ocurrio un error en append " + err.Error())
		return false
	}
	_, err = arch.WriteString(tex)
	if err != nil {
		fmt.Println("Error al escribir el archivo " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() {
	arc, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	fmt.Println(string(arc))

}

func LeerArchivop() {
	arc, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(arc)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println("- " + linea)
	}
	arc.Close()
}
