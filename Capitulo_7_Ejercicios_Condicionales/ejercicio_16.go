package main

import "fmt"

func main() {
	/* 16. Leer un número entero de tres dígitos y determinar si al menos dos de sus tres dígitos son
	iguales.*/
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
		if udig == ddig || udig == cdig || ddig == cdig {
			fmt.Println("Se determina que al menos 2 de los 3 dígitos son iguales ")
		} else {
			fmt.Println("Se determina que el número ingresado no tiene dígitos iguales ")
		}
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
