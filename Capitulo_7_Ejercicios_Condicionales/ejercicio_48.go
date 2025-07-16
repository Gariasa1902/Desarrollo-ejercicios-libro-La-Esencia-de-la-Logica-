package main

import "fmt"

func main() {
	// 48. Leer un número entero y si es menor que 100 determinar si es primo.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num < 100 {
		var i int32 = 1 // CONTROLA EL CICLO WHILE
		var cont int32 = 0
		for i <= num {
			if num%i == 0 {
				cont = cont + 1
			}
			i = i + 1
		}
		if cont == 2 {
			fmt.Println("Se determina que el número ingresado es primo: ", num)
		} else {
			fmt.Println("Se determina que el número ingresado no es primo: ", num)
		}
	}
}
