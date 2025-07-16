package main

import "fmt"

func main() {
	// 9. Leer un número entero de dos dígitos y determinar si un dígito es múltiplo del otro.
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-*1)
	}

	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		if ddig%udig == 0 {
			fmt.Println("El dígito " + ddig + " es multiplo del dígito " + udig)
		} else {
			fmt.Println("El dígito " + ddig + " no es multiplo del dígito " + udig)
		}
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}
}
