package main

import "fmt"

func main() {
	/* 29. Leer un número entero de cinco dígitos y determinar si es un número capicúo. Ej. 15651,
	59895.*/
	var num int32
	fmt.Println("Ingrese un número entero de 5 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10000 && num <= 99999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		//var cdig int32 = (num % 1000) / 100
		var ucdig int32 = (num % 10000) / 1000
		var ducdig int32 = (num % 100000) / 10000

		if udig == ducdig && ddig == ucdig {
			fmt.Println("Se determina que el número ingresado es un número capícua: ", num)
		} else {
			fmt.Println("Se determina que el número ingresado no es un número capícua: ", num)
		}
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}

}
