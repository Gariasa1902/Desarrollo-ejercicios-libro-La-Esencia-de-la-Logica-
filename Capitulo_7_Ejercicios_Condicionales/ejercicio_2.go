package main

import "fmt"

func main() {
	// 2. Leer un número entero y determinar si tiene 3 dígitos.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 100 && num <= 999 {
		fmt.Println("El número ingresado es de tres dígitos ")
	} else {
		fmt.Println("El número ingresado no es de tres dígitos ")
	}
}
