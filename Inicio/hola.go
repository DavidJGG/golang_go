package main

import (
	"fmt"
)

func main() {
	//arreglos()
	//pedirDatos()
	//tipoSlice()
	//punteros()
}

func arreglos() {
	var arreglo [3]string
	var arreglo2 [3]int
	arreglo3 := []int{1, 2, 3}

	fmt.Println(len(arreglo3))

	fmt.Println(arreglo)
	fmt.Println(arreglo2)
	fmt.Println(arreglo3)
}

func pedirDatos() {
	var numTabla, longTabla int
	fmt.Println("Ingrese el numero de la tabla: ")
	fmt.Scanf("%v\n", &numTabla)
	fmt.Println("Ingrese la longitud de la tabla: ")
	fmt.Scanf("%v\n", &longTabla)

	for i := 0; i <= longTabla; i++ {
		fmt.Printf("%v * %v = %v \n", numTabla, i, (numTabla * i))
	}
}

func tipoSlice() {
	var vec []int
	vec2 := []int{1, 2, 3, 4}
	fmt.Println(vec)
	fmt.Println(vec2)

	println("------ SLICING ---------")
	vec3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(vec3)
	miSlice1 := vec3[2:5] //<--- Slicing
	miSlice2 := vec3[:5]  //<--- Slicing
	miSlice3 := vec3[5:8] //<--- Slicing

	fmt.Println(miSlice1)
	fmt.Println(miSlice2)
	fmt.Println(miSlice3)

	println(" ---------------- MAKE Y APEND -------------------")
	slice := make([]int, 3, 55)
	fmt.Println(cap(slice))
	fmt.Println(slice)
	slice = append(slice, 2)
	fmt.Println(slice)
	println(" ---------------- COPY -------------------")
	//copy(destino, fuente)
	sliceCPY1 := []int{1, 2, 3, 4}
	sliceCPY2 := make([]int, 2)
	copy(sliceCPY2, sliceCPY1)
	fmt.Println(sliceCPY1)
	fmt.Println(sliceCPY2)
}

func punteros() {
	var x, y *int
	entero := 5
	x = &entero
	y = &entero
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*x)
	fmt.Println(*y)
}
