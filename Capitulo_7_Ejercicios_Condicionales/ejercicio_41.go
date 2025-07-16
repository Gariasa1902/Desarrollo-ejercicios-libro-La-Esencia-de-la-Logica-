package main

import "fmt"

func main() {
	// 41. Leer dos números enteros y determinar si la diferencia entre los dos es un número primo.
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
	var i int32 = 1 //CONTROLA EL CICLO WHILE
	var cont int32 = 0
	if num1 > num2 {
		dif = num1 - num2
		fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
		for i <= dif {
			if dif%i == 0 {
				cont = cont + 1
			}
			i = i + 1
		}
		if cont == 2 {
			fmt.Println("Se determina que la diferencia entre los 2 numeros ingresaos es un número pimo ")
		} else {
			fmt.Println("Se determina que la diferencia entre los 2 numeros ingresaos no es un número pimo ")
		}
	} else {
		if num2 > num1 {
			dif = num2 - num1
			fmt.Println("La diferente entre los 2 número ingresados es: ", dif)
			for i <= dif {
				if dif%i == 0 {
					cont = cont + 1
				}
				i = i + 1
			}
			if cont == 2 {
				fmt.Println("Se determina que la diferencia entre los 2 numeros ingresaos es un número pimo ")
			} else {
				fmt.Println("Se determina que la diferencia entre los 2 numeros ingresaos no es un número pimo ")
			}
		} else {
			fmt.Println("Los número ingresados son iguales ")
		}
	}

}
