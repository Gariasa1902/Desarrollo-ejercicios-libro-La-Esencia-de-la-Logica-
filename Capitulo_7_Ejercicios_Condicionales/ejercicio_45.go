package main

import "fmt"

func main() {
	/* 45. Leer un número entero de 2 dígitos y si es par mostrar en pantalla la suma de sus dígitos, si es
	primo y menor que 10 mostrar en pantalla su último dígito y si es múltiplo de 5 y menor que 30
	mostrar en pantalla el primer dígito.*/
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}
	var suma int32 = 0
	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		if (num/2)*2 == 0 {
			suma = ddig + udig
			fmt.Println("El número ingresado es par y el resultado de la suma de sus dígitos es: ", suma)
		}
		if suma < 10 {
			if suma == 2 || suma == 3 || suma == 5 || suma == 7 {
				fmt.Println(udig)
			}
		}
		if num%5 == 0 && num < 30 {
			fmt.Println("Se determina que el número ingresado es múltiplo de 5 y es menor que 30 y su primer dígito es: ", ddig)
		}
	} else {
		fmt.Println("El número ingresado debe ser de 2 dígitos ")
	}
}
