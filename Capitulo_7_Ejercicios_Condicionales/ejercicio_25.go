package main

import "fmt"

func main() {
	/* 25. Leer un número entero de tres dígitos y determinar si alguno de sus dígitos es igual a la suma
	de los otros dos.*/
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

		if udig+ddig == cdig {
			fmt.Println("Se determina que el dígito cdig ", cdig, " es igual a la suma de los otros dígitos ", udig, " y ", ddig)
		} else {
			if udig+cdig == ddig {
				fmt.Println("Se determina que el dígito ddig ", ddig, " es igual a la suma de los otros dígitos ", udig, " y ", cdig)
			} else {
				if ddig+cdig == udig {
					fmt.Println("Se determina que el dígito udig ", udig, " es igual a la suma de los otros dígitos ", ddig, " y ", cdig)
				} else {
					fmt.Println("Se determina que el resultados de sumar dos dígitos no da como resultado otro de los dígitos ")
				}
			}
		}
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
