package main

import "fmt"

func main() {
	/* 50. Leer un número entero y si es múltiplo de 4 mostrar en pantalla su mitad, si es múltiplo de 5
	mostrar en pantalla su cuadrado y si es múltiplo e 6 mostrar en pantalla su primer dígito.
	Asumir que el número no es mayor que 100.*/
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num < 100 {
		if num%4 == 0 {
			var mitad int32 = num / 2
			fmt.Println("El número es múltiplo de 4, su mitad es: ", mitad)
		} else {
			if num%5 == 0 {
				var cuadrado int32 = num * num
				fmt.Println("El número es múltiplo de 5, su cuadrado es: ", cuadrado)
			} else {
				if num%6 == 0 {
					var ddig int32 = (num / 10) % 10
					fmt.Println("El número es múltiplo de 6, su primer dígito es: ", ddig)
				}
			}
		}
	} else {
		fmt.Println("El número ingresado es mayor que 100, por favor ingrese un número menor o igual a 100.")
	}
}
