package main

import "fmt"

func main() {
	// 26. Leer un número entero de cuatro dígitos y determinar a cuanto es igual la suma de sus dígitos.
	var num int32
	fmt.Println("Ingrese un número entero de 4 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num - (-1)
	}

	if num >= 1000 && num <= 9999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100
		var ucdig int32 = (num % 10000) / 1000

		var suma int32 = udig + ddig + cdig + ucdig
		fmt.Println("Se determina que el resultado de la suma de los 4 dígitos del número ingresado es: ", suma)
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}
}
