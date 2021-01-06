package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	lsSimp := lse.newLista()
	lsSimp.insertar(1)
	lsSimp.insertar(2)
	lsSimp.insertar(3)
	lsSimp.insertar(4)
	lsSimp.insertar(5)
	lsSimp.insertar(6)
	lsSimp.imprimir()
}
