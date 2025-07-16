package main

import "fmt"

func main() {
	// Leer un número entero y determinar si es positivo o negativo o si es 0.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("El número ingresado es negativo ")
	}
	if num > 0 {
		fmt.Println("El número ingresado es positivo ")
	}
	if num == 0 {
		fmt.Println("El número ingresado es cero ")
	}
}
