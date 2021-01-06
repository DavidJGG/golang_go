package main

import (
	"fmt"
)

//usuario
type User struct {
	edad             int
	nombre, apellido string
}

func (this *User) detalle() string {
	return (" - " + this.nombre + " - " + this.apellido + "\n")
}

func (this *User) setNombre(n string) {
	this.nombre = n
}

func main() {

	per := User{edad: 23, nombre: "David", apellido: "González"}
	fmt.Println("--------- FUNCION DENTRO DE STRUCT ----------")
	fmt.Println(per.detalle())
	per.setNombre("Otro nombre")
	fmt.Println(per.detalle())
	fmt.Println("--------- por valor ----------")
	per2 := per
	fmt.Println(per, per2)
	per2.apellido = "Apellido editado"
	fmt.Println(per, per2)
	fmt.Println("--------- por referencia ----------")
	per1 := User{23, "David", "González"}
	// var per3 *User
	// per3 = &per1 // Go infiere el tipo puntero con :=
	per3 := &per1
	fmt.Println(per1, per3)
	(*per3).apellido = "Editado"
	fmt.Println(per1, per3)

	fmt.Println("----------- CON NEW -----------------")
	per4 := new(User)
	(*per4).nombre = "nombre1"
	(*per4).apellido = "apellido1"
	(*per4).edad = 22
	fmt.Println(per4, *per4)
	per5 := per4
	per5.apellido = "sss"
	fmt.Println(per5, per4)
}
