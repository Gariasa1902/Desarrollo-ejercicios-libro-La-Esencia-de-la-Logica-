package main

import "fmt"

func main() {
	/* 46. Leer un número entero de 2 dígitos y si terminar en 1 mostrar en pantalla su primer dígito, si
	termina en 2 mostrar en pantalla la suma de sus dígitos y si termina en 3 mostrar en pantalla el
	producto de sus dos dígitos.*/
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}
	var udig int32 = num % 10
	var ddig int32 = (num % 100) / 10
	switch udig {
	case 1:
		if udig == 1 {
			fmt.Println(ddig)
		}
	case 2:
		if udig == 2 {
			fmt.Println(ddig + udig)
		}
	case 3:
		if udig == 3 {
			fmt.Println(ddig * udig)
		}
	}
}
