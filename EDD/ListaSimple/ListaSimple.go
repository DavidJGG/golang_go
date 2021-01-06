package ListaSimple

import "fmt"

type Nodo struct {
	valor     int
	siguiente *Nodo
}

func (this *Nodo) ImprimirNodo() {
	fmt.Println(*this)
	if this.siguiente == nil {
		return
	}
	(*this).siguiente.ImprimirNodo()
}

func (this *Nodo) getValor() int {
	return this.valor
}

//*********************************************

type Lista struct {
	primero *Nodo
	ultimo  *Nodo
}

func NewLista() Lista {
	return Lista{nil, nil}
}

func (this *Lista) GetPrimero() *Nodo {
	return this.primero
}

func (this Lista) nuevoNodo(val int) *Nodo {
	nuevo := new(Nodo)
	(*nuevo).valor = val
	(*nuevo).siguiente = nil
	return nuevo
}

func (this *Lista) Insertar(val int) {
	nuevo := this.nuevoNodo(val)
	if this.primero == nil {
		this.primero = nuevo
		this.ultimo = nuevo
	} else {
		this.ultimo.siguiente = nuevo
		this.ultimo = nuevo
	}
}

func (this *Lista) Imprimir() {
	fmt.Println(this.primero)
	aux := this.primero
	for aux != nil {
		fmt.Printf("[ %v ] -> ", aux.valor)
		aux = aux.siguiente
	}
	fmt.Println(" null")
}
