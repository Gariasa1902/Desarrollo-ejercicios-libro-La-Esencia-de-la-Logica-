package main

import "fmt"

func main() {
	// 15. Leer un número entero de tres dígitos y determinar a cuánto es igual la suma de sus dígitos.
	var num int32
	fmt.Println("Ingrese un número entero de 3 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 100 && num <= 999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100

		var suma int32 = udig + ddig + cdig
		fmt.Println("Se determina que el resultado de la suma de los 3 dígitos es ", suma)
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
