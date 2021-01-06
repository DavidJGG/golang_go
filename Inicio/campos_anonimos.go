package main

import "fmt"

type Humano struct {
	nombre string
}

type Tutor struct {
	Humano // <- campo anónimo
	puesto string
}

func main() {
	tuto := Tutor{Humano{"David"}, "Profesor"}
	fmt.Println(tuto)
}
