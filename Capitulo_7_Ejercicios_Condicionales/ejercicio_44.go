package main

import "fmt"

func main() {
	/* 44. Leer un número entero de 4 dígitos y determinar si el primer dígito es múltiplo de alguno de los
	otros dígitos*/
	var num int32
	fmt.Println("Ingrese un número entero de 4 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}
	if num >= 1000 && num <= 9999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100
		var ucdig int32 = (num % 10000) / 1000
		if ucdig%udig == 0 {
			fmt.Println("Se determina que el primer digito del número ingresado: ", ucdig, " es múltiplo del último dígito: ", udig)
		} else {
			if ucdig%ddig == 0 {
				fmt.Println("Se determina que el primer digito del número ingresado: ", ucdig, " es múltiplo del peúltimo dígito: ", ddig)
			} else {
				if ucdig%cdig == 0 {
					fmt.Println("Se determina que el primer digito del número ingresado: ", ucdig, " es múltiplo del antepeúltimo dígito: ", cdig)
				} else {
					fmt.Println("Se determina que el primer digito del número ingresado: ", ucdig, " no es múltiplo de ninguno de los otros dígitos ")
				}
			}
		}
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}
}
