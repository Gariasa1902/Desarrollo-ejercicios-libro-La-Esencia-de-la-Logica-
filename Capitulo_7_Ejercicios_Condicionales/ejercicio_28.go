package main

import "fmt"

func main() {
	// 28. Leer un número entero menor que 50 y positivo y determinar si es un número primo.
	var num int32
	fmt.Println("Ingrese un número entero menor que 50 ")
	fmt.Scan(&num)

	if num < 0 {
		num = num - (-1)
	}

	if num > 0 {
		if num < 50 {
			if num == 2 || num == 3 || num == 5 || num == 7 || num == 11 || num == 13 || num == 17 || num == 19 ||
				num == 23 || num == 29 || num == 31 || num == 37 || num == 41 || num == 43 || num == 47 {
				fmt.Println("Se determina que el número ingresado es menor que 50 y primo ")
			} else {
				fmt.Println("Se determina que el número ingresado es menor que 50 y no es primo ")
			}
		} else {
			fmt.Println("El número ingresado debe ser menor que 50 ")
		}
	} else {
		fmt.Println("El número ingresado debe ser positivo ")
	}

}
