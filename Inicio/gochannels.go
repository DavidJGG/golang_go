package main

import (
	"fmt"
	"time"
)

func main() {
	// channel := make(chan string) //Que se transfiere entre canales
	// go func(channel chan string) {
	// 	for {
	// 		var nombre string
	// 		fmt.Scanln(&nombre)
	// 		channel <- nombre
	// 	}
	// }(channel)

	// msg := <-channel
	// fmt.Println("El canal dice: " + msg)

	ejemplo()

}

func ejemplo() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go func(canal1 chan int) {
		i := 0
		for true {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 20)
				canal1 <- i
				//fmt.Println(i)
			}
			i++
		}

	}(canal1)

	go impares(canal2)

	for true {
		txt_p := <-canal1
		fmt.Println(txt_p)
		txt_i := <-canal2
		fmt.Println(txt_i)
	}
	// ss := ""
	// fmt.Scanln(&ss)
}

func impares(canal2 chan int) {
	i := 0
	for true {
		if i%2 != 0 {
			time.Sleep(time.Millisecond * 5000)
			canal2 <- i
			//fmt.Println(i)
		}
		i++
	}
}
