package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go separar_nombre("David")
	go func() {
		fmt.Println("Go rutina func anónima")
	}()
	var wait string
	fmt.Scanln(&wait)
}

func separar_nombre(nombre string) {
	letras := strings.Split(nombre, "")
	for index, letra := range letras {
		fmt.Printf("[%v]: %v \n", index, letra)
		time.Sleep(500 * time.Millisecond)
	}
}
