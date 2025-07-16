package main

import "fmt"

func main() {
	/* 2. Leer dos números enteros y mostrar todos los enteros
	comprendidos entre el menor y el mayor*/
	var num1, num2 int32
	fmt.Println("Ingrese el primer número entero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo número entero: ")
	fmt.Scan(&num2)
	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1) // Convertir a positivo si es negativo
		num2 = num2 * (-1) // Convertir a positivo si es negativo
	}
	if num1 < num2 {
		for num1 <= num2 {
			fmt.Println(num1)
			num1++
		}
	} else {
		if num2 < num1 {
			for num2 <= num1 {
				fmt.Println(num2)
				num2++
			}
		} else {
			fmt.Println("Se determina que los 2 números ingresados son iguales ")
		}
	}
}
