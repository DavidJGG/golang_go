package lse

import "fmt"

type Nodo struct {
	valor     int
	siguiente *Nodo
}

func (this *Nodo) imprimirNodo() {
	fmt.Println(this)
}

func (this *Nodo) getValor() int {
	return this.valor
}

//*********************************************

type Lista struct {
	primero *Nodo
	ultimo  *Nodo
}

func (this Lista) nuevoNodo(val int) *Nodo {
	nuevo := new(Nodo)
	(*nuevo).valor = val
	(*nuevo).siguiente = nil
	return nuevo
}

func (this Lista) insertar(val int) {
	nuevo := this.nuevoNodo(val)
	if this.primero == nil {
		this.primero = nuevo
		this.ultimo = nuevo
	} else {
		this.ultimo.siguiente = nuevo
		this.ultimo = nuevo
	}
}

func (this Lista) imprimir() {
	aux := this.primero
	for aux != nil {
		fmt.Printf("[ %v ] -> ", aux.valor)
		aux = aux.siguiente
	}
	fmt.Println(" null")
}

func newLista() Lista {
	return Lista{nil, nil}
}
