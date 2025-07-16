package main

import "fmt"

func main() {
	// 38. Leer tres números enteros y determinar si el último dígito de los tres números es igual.
	var num1 int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero ")
	fmt.Scan(&num2)
	var num3 int32
	fmt.Println("Ingrese un último número entero ")
	fmt.Scan(&num3)

	if num1 < 0 && num2 < 0 && num3 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
		num3 = num3 * (-1)
	}

	var udig1 int32 = num1 % 10
	var udig2 int32 = num2 % 10
	var udig3 int32 = num3 % 10
	if udig1 == udig2 && udig1 == udig3 && udig2 == udig3 {
		fmt.Println("Se determina que el último dígito de los 3 números ingresados son iguales ")
	} else {
		fmt.Println("Se determina que el último dígito de los 3 números no ingresados son iguales ")
	}
}
