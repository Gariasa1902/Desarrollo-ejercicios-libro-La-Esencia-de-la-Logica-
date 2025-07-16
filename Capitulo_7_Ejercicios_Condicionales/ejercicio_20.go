package main

import "fmt"

func main() {
	// 20. Leer tres números enteros y mostrarlos ascendentemente.
	var num1 int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero ")
	fmt.Scan(&num2)
	var num3 int32
	fmt.Println("Ingrese un último número entero ")
	fmt.Scan(&num3)

	if num1 < 0 && num2 < 0 && num3 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
		num3 = num3 * (-1)
	}

	if num1 > num2 && num2 > num3 {
		fmt.Println(num1)
		fmt.Println(num2)
		fmt.Println(num3)
	} else {
		if num2 > num1 && num1 > num3 {
			fmt.Println(num2)
			fmt.Println(num1)
			fmt.Println(num3)
		} else {
			if num3 > num2 && num2 > num1 {
				fmt.Println(num3)
				fmt.Println(num2)
				fmt.Println(num1)
			} else {
				fmt.Println("Se determina que todos los número son iguales ")
			}
		}
	}
}
