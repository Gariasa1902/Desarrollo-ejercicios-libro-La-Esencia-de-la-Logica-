package main

import "fmt"

func main() {
	// 32. Leer un número entero y determinar si es múltiplo de 7.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num%7 == 0 {
		fmt.Println("Se determina que el número ingresado: ", num, " es múltiplo de 7 ")
	} else {
		fmt.Println("Se determina que el número ingresado: ", num, " no es múltiplo de 7 ")
	}

}
