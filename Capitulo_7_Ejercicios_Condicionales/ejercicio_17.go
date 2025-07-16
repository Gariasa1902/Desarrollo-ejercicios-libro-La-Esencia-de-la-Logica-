package main

import "fmt"

func main() {
	// 17. Leer un número entero de tres dígitos y determinar en qué posición está el mayor dígito.
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

		if udig > ddig && udig > cdig {
			fmt.Println("Se determina que el dígito mayor esta ubicado en la posición unidad: ", udig)
		} else {
			if ddig > udig && ddig > cdig {
				fmt.Println("Se determina que el dígito mayor esta ubicado en la posición decena: ", ddig)
			} else {
				if cdig > udig && cdig > ddig {
					fmt.Println("Se determina que el dígito mayor esta ubicado en la posición centena: ", cdig)
				} else {
					fmt.Println("Se determina que todos los dígitos son iguales")
				}
			}
		}
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
