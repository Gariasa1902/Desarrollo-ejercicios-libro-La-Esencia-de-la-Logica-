package main

import "fmt"

func main() {
	// 22. Leer un número entero de tres dígitos y determinar si el primer dígito es igual al último.
	var num int32
	fmt.Println("Ingrese un número entero de 3 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 100 && num <= 999 {
		var udig int32 = num % 10
		//var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100

		if udig == cdig {
			fmt.Println("Se determina que el primer dígito: ", cdig, " es igual al último dígito: ", udig)
		} else {
			fmt.Println("Se determina que el número ingresado no tiene dígitos iguales ")
		}
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
