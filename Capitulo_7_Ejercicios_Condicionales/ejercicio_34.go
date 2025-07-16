package main

import "fmt"

func main() {
	// 34. Leer un número entero menor que mil y determinar cuántos dígitos tiene.
	var num int32
	fmt.Println("Ingrese un número entero menor que 1000 ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num < 1000 {
		if num < 10 {
			fmt.Println("Se determina que el número ingresado tiene 1 dígito ")
		}
		if num >= 10 && num <= 99 {
			fmt.Println("Se determina que el número ingresado tiene 2 dígito ")
		}
		if num >= 100 && num <= 999 {
			fmt.Println("Se determina que el número ingresado tiene 3 dígito ")
		}
	} else {
		fmt.Println("El número ingresado debe ser menor a 1000 ")
	}

}
