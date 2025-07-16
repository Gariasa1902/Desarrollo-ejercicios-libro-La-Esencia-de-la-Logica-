package main

import "fmt"

func main() {
	// 7. Leer un número entero de dos dígitos y determinar si es primo y además si es negativo.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		if num >= -99 && num <= -10 {
			if num == -11 || num == -13 || num == -17 || num == -19 || num == -23 ||
				num == -29 || num == -31 || num == -37 || num == -41 || num == -43 ||
				num == -47 || num == -53 || num == -59 || num == -61 || num == -67 ||
				num == -71 || num == -73 || num == -79 || num == -83 || num == -89 || num == -97 {
				fmt.Println("El número ingresado es primo y ademas es negativo ")
			} else {
				fmt.Println("El número ingresado no es primo y ademas es negativo ")
			}
		} else {
			fmt.Println("El número ingresado no es de 2 dígitos ")
		}
	} else {
		if num >= 10 && num <= 99 {
			if num == 11 || num == 13 || num == 17 || num == 19 || num == 23 ||
				num == 29 || num == 31 || num == 37 || num == 41 || num == 43 ||
				num == 47 || num == 53 || num == 59 || num == 61 || num == 67 ||
				num == 71 || num == 73 || num == 79 || num == 83 || num == 89 || num == 97 {
				fmt.Println("El número ingresado no es negativo, pero es un número primo ")
			}
		} else {
			fmt.Println("El número ingresado no es negativo y no es un número primo ")
		}
	}
}
