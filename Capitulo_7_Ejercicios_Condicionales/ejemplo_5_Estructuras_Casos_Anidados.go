package main

import "fmt"

func main() {
	/* Leer un entero y si es igual a cualquier dígito comprendido entre 1 y 5 escribir su nombre. Si es
	igual a cinco además de escribir su nombre leer otro dígito y, si este último está entre 1 y 5, escribir
	su componente decimal. Si entró un 3 entonces escribir “Cincuenta y Tres”, si entró un 1 entonces
	escribir “Cincuenta y Uno”.*/
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
		var dig int32
		fmt.Println("Ingrese otro número entero ")
		fmt.Scan(&dig)
		switch dig {
		case 1:
			dig = 1
			fmt.Println("Cincuenta y Uno")
		case 2:
			dig = 2
			fmt.Println("Cincuenta y Dos")
		case 3:
			dig = 3
			fmt.Println("Cincuenta y Tres")
		case 4:
			dig = 4
			fmt.Println("Cincuenta y Cuatro")
		case 5:
			dig = 5
			fmt.Println("Cincuenta y Cinco")
		default:
			fmt.Println("El número es mayor que cinco")
		}
	default:
		fmt.Println("El número es mayor que cinco")
	}
}
