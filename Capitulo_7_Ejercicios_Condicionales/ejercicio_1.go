package main

import "fmt"

func main() {
	// 1. Leer un número entero y determinar si es un número terminado en 4.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	var ud int32
	ud = num % 10
	if ud == 4 {
		fmt.Println("El número ingresado termina en 4 ")
	} else {
		fmt.Println("El número ingresado no termina en 4 ")
	}

}
