package main

import "fmt"

func main() {
	// 30. Leer un número entero de cuatro dígitos y determinar si el segundo dígito es igual al penúltimo.
	var num int32
	fmt.Println("Ingrese un número entero de 4 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}
	if num >= 1000 && num <= 9999 {
		//var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100
		//var ucdig int32 = (num % 10000) / 1000
		if ddig == cdig {
			fmt.Println("Se determina que el segundo dígito: ", ddig, " es igual al penúltimo dígito: ", cdig)
		} else {
			fmt.Println("Se determina que el segundo dígito: ", ddig, " no es igual al penúltimo dígito: ", cdig)
		}
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}
}
