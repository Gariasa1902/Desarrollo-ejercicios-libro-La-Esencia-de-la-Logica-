package main

import "fmt"

func main() {
	// Desarrollar un programa que permite leer un número entero positivo y determinar si es par.
	var num int16
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("El número ingresado no sirve para nuestro propósito ")
	} else {
		if (num/2)*2 == num {
			fmt.Println("El número ingresado es par ")
		} else {
			fmt.Println("El número ingresado no es par ")
		}
	}

}
