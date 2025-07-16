package main

import "fmt"

func main() {
	/* 43. Leer dos números enteros y determinar si la diferencia entre los dos es un número divisor
	exacto de alguno de los dos números.*/
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
		fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
		if num1%dif == 0 {
			fmt.Println("Se determina que la diferencia entre los 2 número ingresados: ", dif, " es divisor exacto del primer número ingresado: ", num1)
		} else {
			if num2%dif == 0 {
				fmt.Println("Se determina que la diferencia entre los 2 número ingresados: ", dif, " es divisor exacto del segundo número ingresado: ", num2)
			} else {
				fmt.Println("Se determina que la diferencia de los 2 números ingresasos no es divisor exacto de ninguno de los números ")
			}
		}
	} else {
		if num2 > num1 {
			dif = num2 - num1
			fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
			if num1%dif == 0 {
				fmt.Println("Se determina que la diferencia entre los 2 número ingresados: ", dif, " es divisor exacto del primer número ingresado: ", num1)
			} else {
				if num2%dif == 0 {
					fmt.Println("Se determina que la diferencia entre los 2 número ingresados: ", dif, " es divisor exacto del segundo número ingresado: ", num2)
				} else {
					fmt.Println("Se determina que la diferencia de los 2 números ingresasos no es divisor exacto de ninguno de los números ")
				}
			}
		} else {
			fmt.Println("Los número ingresados son iguales ")
		}
	}
}
