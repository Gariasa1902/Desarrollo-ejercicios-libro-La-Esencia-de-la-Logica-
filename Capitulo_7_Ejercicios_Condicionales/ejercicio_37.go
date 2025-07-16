package main

import "fmt"

func main() {
	// 37. Leer dos números enteros y determinar cuál es múltiplo de cuál.
	var num1 int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero ")
	fmt.Scan(&num2)

	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
	}

	if num1%num2 == 0 {
		fmt.Println("Se determina que el primer número ingresado ", num1, " es múltiplo del segundo número ingresado ", num2)
	} else {
		if num2%num1 == 0 {
			fmt.Println("Se determina que el segundo número ingresado ", num2, " es múltiplo del primer número ingresado ", num1)
		}
	}
}
