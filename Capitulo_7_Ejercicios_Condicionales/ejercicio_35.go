package main

import "fmt"

func main() {
	/* 35. Leer un número entero de dos dígitos, guardar cada dígito en una variable diferente y luego
	mostrarlas en pantalla.*/
	var num int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 10 && num <= 99 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		fmt.Println(ddig, "y", udig)
	} else {
		fmt.Println("El número ingresado no es de 2 dígitos ")
	}

}
