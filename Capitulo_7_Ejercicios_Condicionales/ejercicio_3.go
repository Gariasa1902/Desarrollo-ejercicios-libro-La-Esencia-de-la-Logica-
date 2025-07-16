package main

import "fmt"

func main() {
	// 3. Leer un número entero y determinar si es negativo.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("El número ingresado es negativo ")
	} else {
		fmt.Println("El número ingresado no es negativo ")
	}
}
