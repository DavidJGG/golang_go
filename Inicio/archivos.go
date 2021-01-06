package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	datos, err := ioutil.ReadFile("./lectura.txt")
	if datos == nil {
		fmt.Println("Error (dato==nil) " + err.Error())
		return
	}
	fmt.Println(string(datos))
}
