package main

import "fmt"

func main() {
	// 5. Leer un número entero de dos dígitos y determinar si ambos dígitos son pares.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		if (udig/2)*2 == udig && (ddig/2)*2 == ddig {
			fmt.Println("Se determina que los 2 dígitos del número ingresado son pares ")
		} else {
			fmt.Println("Se determina que solo uno o ninguno de los 2 dígitos son pares ")
		}
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}

}
