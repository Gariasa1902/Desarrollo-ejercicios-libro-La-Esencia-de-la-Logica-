package main

import "fmt"

func main() {
	// 31. Leer un número entero y determina si es igual a 10.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num == 10 {
		fmt.Println("Se determmina que el número ingresado: ", num, " es igual a 10 ")
	} else {
		fmt.Println("Se determmina que el número ingresado: ", num, " no es igual a 10 ")
	}
}
