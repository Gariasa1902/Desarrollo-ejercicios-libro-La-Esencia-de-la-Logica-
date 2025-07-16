package main

import "fmt"

func main() {
	// 8. Leer un número entero de dos dígitos y determinar si sus dos dígitos son primos.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cont int32 = 0
		if ddig == 2 || ddig == 3 || ddig == 5 || ddig == 7 {
			cont++
		}
		if udig == 2 || udig == 3 || udig == 5 || udig == 7 {
			cont++
		}
		if cont == 2 {
			fmt.Println("Se determina que los 2 dígitios son primos ")
		} else {
			fmt.Println("Solo uno o ninguno de los 2 dígitos es primo ")
		}
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}
}
