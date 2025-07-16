package main

import "fmt"

func main() {
	// 42. Leer dos números enteros y determinar si la diferencia entre los dos es un número par.
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
		if (dif/2)*2 == dif {
			fmt.Println("Se determina que la diferencia entre los 2 número ingresados es un número par ")
		} else {
			fmt.Println("Se determina que la diferencia entre los 2 número ingresados no es un número par ")
		}
	} else {
		if num2 > num1 {
			dif = num2 - num1
			fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
			if (dif/2)*2 == dif {
				fmt.Println("Se determina que la diferencia entre los 2 número ingresados es un número par ")
			} else {
				fmt.Println("Se determina que la diferencia entre los 2 número ingresados no es un número par ")
			}
		} else {
			fmt.Println("Los número ingresados son iguales ")
		}
	}
}
