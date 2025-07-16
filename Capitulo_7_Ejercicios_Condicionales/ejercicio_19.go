package main

import "fmt"

func main() {
	// 19. Leer tres números enteros y determinar cuál es el mayor. Usar solamente dos variables.
	var num1 int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero ")
	fmt.Scan(&num2)

	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
	}

	if num1 > num2 {
		fmt.Println("Ingrese otro número entero ")
		fmt.Scan(&num2)
		fmt.Println("Se determina que el número: ", num1, " es mayor ")
	} else {
		fmt.Println("Ingrese otro número entero ")
		fmt.Scan(&num1)
		if num1 > num2 {
			fmt.Println("Se determina que el número: ", num1, " es mayor ")
		} else {
			if num2 > num1 {
				fmt.Println("Se determina que el número: ", num2, " es mayor ")
			} else {
				fmt.Println("Se determina que los 3 números son iguales ")
			}
		}
	}
}
