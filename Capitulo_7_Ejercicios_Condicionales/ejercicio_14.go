package main

import "fmt"

func main() {
	/* 14. Leer dos números enteros de dos dígitos y determinar a cuánto es igual la suma de todos los
	dígitos.*/
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
		var udig1 int32 = num1 % 10
		var ddig1 int32 = (num1 % 100) / 10

		var udig2 int32 = num2 % 10
		var ddig2 int32 = (num2 % 100) / 10

		var suma int32 = udig1 + ddig1 + udig2 + ddig2
		fmt.Println("Se determina que el resultado de la suma de todos los dígitos es ", suma)
	} else {
		fmt.Println("Solo 1 o ninguno de los 2 numeros es de 2 dígitos ")
	}
}
