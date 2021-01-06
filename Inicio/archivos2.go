package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leer_archivo()
	fmt.Println("Recover ")
}

func leer_archivo() {
	archivo, err := os.Open("./lectura.txtsdssssd")
	defer func() {
		ss := recover() // se recupera del error
		fmt.Println(ss)
		archivo.Close() // <-- se ejecuta cuanda la función se termine,
		//ya sea por errores o por return
	}()

	if err != nil {
		panic(err) // <-- Termina la ejecucion del programa
		//y detalla la linea del error
		fmt.Println("Errores: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	numLn := 1
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Printf("[%v]: %v\n", numLn, linea)
		numLn++
	}

	archivo.Close()
}
