package main

import "fmt"

func main() {
	// 39. Leer tres números enteros y determina si el penúltimo dígito de los tres números es igual.
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

	var ddig1 int32 = (num1 % 100) / 10
	var ddig2 int32 = (num2 % 100) / 10
	var ddig3 int32 = (num3 % 100) / 10
	if ddig1 == ddig2 && ddig1 == ddig3 && ddig2 == ddig3 {
		fmt.Println("Se determina que el peúltimo dígito de los 3 números ingresados son iguales ")
	} else {
		fmt.Println("Se determina que el peúltimo dígito de los 3 números no ingresados son iguales ")
	}
}
