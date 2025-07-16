package main

import "fmt"

func main() {
	/* 13. Leer dos números enteros de dos dígitos y determinar si la suma de los dos números origina
	un número par.*/
	var num1 int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero de 2 dígitos ")
	fmt.Scan(&num2)

	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
	}

	if num1 >= 10 && num1 <= 99 && num2 >= 10 && num2 <= 99 {
		var suma int32 = num1 + num2
		if suma%2 == 0 {
			fmt.Println("Se determina que la suma de los 2 números es par ", suma)
		} else {
			fmt.Println("Se determina que la suma de los 2 números no es par ", suma)
		}
	} else {
		fmt.Println("Solo 1 o ninguno de los 2 numeros es de 2 dígitos ")
	}

}
