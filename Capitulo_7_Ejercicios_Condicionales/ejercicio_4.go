package main

import "fmt"

func main() {
	// 4. Leer un número entero de dos dígitos y determinar a cuánto es igual la suma de sus dígitos.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num <= 99 {
		var udig int32
		var ddig int32
		var suma int32
		udig = num % 10
		ddig = (num % 100) / 10
		suma = udig + ddig
		fmt.Println("El resultado de sumar los dos dígitos del número ingresado es: ", suma)
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}
}
