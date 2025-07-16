package main

import "fmt"

func main() {
	// 6. Leer un número entero de dos dígitos menor que 20 y determinar si es primo.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos menor que 20 ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num < 20 {
		if num == 11 || num == 13 || num == 17 || num == 19 {
			fmt.Println("El número ingresado es primo ")
		} else {
			fmt.Println("El número no ingresado es primo ")
		}
	} else {
		fmt.Println("El número ingresado no es de 2 dígito o es mayor que 20 ")
	}
}
