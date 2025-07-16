package main

import "fmt"

func main() {
	// 49. Leer un número entero y si es múltiplo de 4 determinar si su último dígito es primo.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num%4 == 0 {
		var udig = num % 10
		var i int32 = 1 // CONTROLA EL CICLO WHILE
		var cont int32 = 0
		for i <= udig {
			if udig%i == 0 {
				cont = cont + 1
			}
			i = i + 1
		}
		if cont == 2 {
			fmt.Println("Se determina que el último dígito del número ingresado es primo: ", udig)
		} else {
			fmt.Println("Se determina que el último dígito del número ingresado no es primo: ", udig)
		}
	} else {
		fmt.Println("El número ingresado no es múltiplo de 4: ", num)
	}

}
