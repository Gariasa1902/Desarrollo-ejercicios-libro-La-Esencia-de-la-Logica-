package main

import "fmt"

func main() {
	// 10. Leer un número entero de dos dígitos y determinar si los dos dígitos son iguales.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		if udig == ddig {
			fmt.Println("Los 2 dígitos del número ingresado son iguales ")
		} else {
			fmt.Println("Los 2 dígitos del número ingresado no son iguales ")
		}
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}
}
