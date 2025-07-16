package main

import "fmt"

func main() {
	/* Leer un número entero y determinar si es de uno o dos o tres o cuatro dígitos. Validar que el
	número no sea negativo.*/
	var num int32
	fmt.Println("Ingrese un número entero")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num <= 9 {
		fmt.Println("El número ingresado es de un dígito ")
	} else {
		if num >= 10 && num <= 99 {
			fmt.Println("El número ingresado es de dos dígito ")
		} else {
			if num >= 100 && num <= 999 {
				fmt.Println("El número ingresado es de cuatro dígito ")
			} else {
				if num >= 1000 && num <= 9999 {
					fmt.Println("El número ingresado es de cuatro dígito ")
				} else {
					fmt.Println("El número ingresado tiene más de cuatro dígito ")
				}
			}
		}
	}
}
