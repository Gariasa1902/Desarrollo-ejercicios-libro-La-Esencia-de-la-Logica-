package main

import "fmt"

func main() {
	/* Leer un número entero y determinar si es positivo o negativo. Si es positivo determinar si tiene dos
	dígitos y si es negativo determinar si tiene tres dígitos. Asumir que no puede entrar el número cero.*/
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num > 0 {
		if num >= 10 && num <= 99 {
			fmt.Println("El número ingresado es positivo y tiene dos dígitos ")
		} else {
			fmt.Println("El número ingresado es positivo y no tiene dos dígitos ")
		}
	} else {
		if num < 0 {
			if num >= -999 && num <= -100 {
				fmt.Println("El número ingresado es negativo y tiene tres dígitos ")
			} else {
				fmt.Println("El número ingresado es negativo y no tiene tres dígitos ")
			}
		}
	}
}
