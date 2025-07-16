package main

import "fmt"

func main() {
	// 33. Leer un número entero y determinar si termina en 7.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	var udig int32 = num % 10
	if udig == 7 {
		fmt.Println("Se determina que el número ingresado: ", num, " termina en 7 ")
	} else {
		fmt.Println("Se determina que el número ingresado: ", num, " no termina en 7 ")
	}

}
