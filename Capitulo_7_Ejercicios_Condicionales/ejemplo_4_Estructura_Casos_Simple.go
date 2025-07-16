package main

import "fmt"

func main() {
	/* Leer un número entero y si es de un dígito y además es menor que 5 escribir su nombre en
	pantalla (El nombre del 1 es UNO, el nombre del 2 es DOS, etc.).*/
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	switch num {
	case 1:
		num = 1
		fmt.Println("Uno")
	case 2:
		num = 2
		fmt.Println("Dos")
	case 3:
		num = 3
		fmt.Println("Tres")
	case 4:
		num = 4
		fmt.Println("Cuatro")
	case 5:
		num = 5
		fmt.Println("Cinco")
	default:
		fmt.Println("El número es mayor que cinco")
	}
}
