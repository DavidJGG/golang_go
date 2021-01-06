package main

import (
	"./ListaSimple"
)

func main() {
	lsSimp := ListaSimple.NewLista()
	lsSimp.Insertar(1)
	lsSimp.Insertar(2)
	lsSimp.Insertar(3)
	lsSimp.Insertar(4)
	lsSimp.Insertar(5)
	lsSimp.Insertar(6)
	lsSimp.Imprimir()
	primero := lsSimp.GetPrimero()
	primero.ImprimirNodo()
}
