package main

import "fmt"

func main() {
	/* Leer dos números enteros positivos y determinar si el último dígito de un número es igual al último
	dígito del otro.*/
	var num1 int16
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num1)
	var num2 int16
	fmt.Println("ingrese otro número entero de 2 dígitos ")
	fmt.Scan(&num2)

	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
	}

	var ud1 int16 = (num1 / 10) * 10
	var ud2 int16 = (num2 / 10) * 10

	if ud1 == ud2 {
		fmt.Println("El último dígito de un número es igual al último dígito del otro. ")
	} else {
		fmt.Println("El último dígito de un número no es igual al último dígito del otro. ")
	}
}
