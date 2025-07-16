package main

import "fmt"

func main() {
	/* 40. Leer dos números enteros y si la diferencia entre los dos es menor o igual a 10 entonces
	mostrar en pantalla todos los enteros comprendidos entre el menor y el mayor de los números
	leídos.*/
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

	var dif int32 = 0
	if num1 > num2 {
		dif = num1 - num2
		if dif <= 10 {
			fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
			for num2 < num1 {
				num2 = num2 + 1
				fmt.Println(num2)
			}
		}
	} else {
		if num2 > num1 {
			dif = num2 - num1
			if dif <= 10 {
				fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
				for num1 < num2 {
					num1 = num1 + 1
					fmt.Println(num1)
				}
			}
		} else {
			fmt.Println("Los número ingresados son iguales ")
		}
	}

}
