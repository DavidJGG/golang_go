package main

import "fmt"

type Usuario interface {
	Permisos() int
	Detalle() string
}

type Persona struct {
	permiso int
	detalle string
}

type Trabajador struct {
	permiso int
	nombre  string
}

func (this Persona) Permisos() int {
	return 5 + this.permiso
}

func (this Persona) Detalle() string {
	return "El detalle es: " + this.detalle
}

func (this Trabajador) Permisos() int {
	return 5 + this.permiso
}

func (this Trabajador) Detalle() string {
	return "El detalle es: Nombre = " + this.nombre
}

func que_es(usu Usuario) string {
	if usu.Permisos() > 10 {
		return "Es admin"
	} else {
		return "No es admin"
	}
}

func main() {
	fmt.Println("------ Interfaces en go ---------")
	per := Persona{5, " mi detalle"}
	fmt.Println(per.Permisos())
	fmt.Println(per.Detalle())
	fmt.Println(que_es(per))
	per.permiso = 11
	fmt.Println(que_es(per))
	fmt.Println("---------- Array de usuarios -------")
	per1 := Persona{5, " detalle de per 1"}
	trab1 := Trabajador{6, " Piterson Ji"}
	usuarios := []Usuario{per1, trab1}
	fmt.Println(usuarios)
	for _, usuario := range usuarios {
		println(usuario.Detalle())
	}
}
