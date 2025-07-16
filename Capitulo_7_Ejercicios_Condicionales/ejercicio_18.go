package main

import "fmt"

func main() {
	// 18. Leer un número entero de tres dígitos y determinar si algún dígito es múltiplo de los otros.
	var num int32
	fmt.Println("Ingrese un número entero de 3 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 100 && num <= 999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100

		if udig%ddig == 0 && udig%cdig == 0 {
			fmt.Println("Se determina que el dígito unidad: ", udig, " es múltiplo de los otros 2 dígitos ")
		} else {
			if ddig%udig == 0 && ddig%cdig == 0 {
				fmt.Println("Se determina que el dígito decena: ", ddig, " es múltiplo de los otros 2 dígitos ")
			} else {
				if cdig%udig == 0 && cdig%ddig == 0 {
					fmt.Println("Se determina que el dígito centena: ", cdig, " es múltiplo de los otros 2 dígitos ")
				} else {
					System.out.println("Se determina que ninguno de los 3 dígitos es multiplo del otro ")
				}
			}
		}
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
