package main

import "fmt"

func main() {
	/* 21. Leer tres números enteros de dos dígitos cada uno y determinar en cuál de ellos se encuentra
	el mayor dígito.*/
	var num1 int32
	fmt.Println("Ingrese un número entero de 2 dígitos ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero de 2 dígitos ")
	fmt.Scan(&num2)
	var num3 int32
	fmt.Println("Ingrese un último número entero de 2 dígitos ")
	fmt.Scan(&num3)

	if num1 < 0 && num2 < 0 && num3 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
		num3 = num3 * (-1)
	}

	if num1 >= 10 && num1 <= 99 && num2 >= 10 && num2 <= 99 && num3 >= 10 && num3 <= 99 {
		var udig1 int32 = num1 % 10
		var ddig1 int32 = (num1 % 100) / 10
		var udig2 int32 = num2 % 10
		var ddig2 int32 = (num2 % 100) / 10
		var udig3 int32 = num3 % 10
		var ddig3 int32 = (num3 % 100) / 10

		if udig1 > ddig1 && udig1 > udig2 && udig1 > ddig2 && udig1 > udig3 && udig1 > ddig3 {
			fmt.Println("Se determina que el dígito mayor: ", udig1, " se encuentra en número num1: ", num1)
		} else {
			if ddig1 > udig1 && ddig1 > udig2 && ddig1 > ddig2 && ddig1 > udig3 && ddig1 > ddig3 {
				fmt.Println("Se determina que el dígito mayor: ", ddig1, " se encuentra en número num1: ", num1)
			} else {
				if udig2 > ddig2 && udig2 > udig1 && udig2 > ddig1 && udig2 > udig3 && udig2 > ddig3 {
					fmt.Println("Se determina que el dígito mayor: ", udig2, " se encuentra en número num2: ", num2)
				} else {
					if ddig2 > udig2 && ddig2 > udig1 && ddig2 > ddig1 && ddig2 > udig3 && ddig2 > ddig3 {
						fmt.Println("Se determina que el dígito mayor: ", ddig2, " se encuentra en número num2: ", num2)
					} else {
						if udig3 > ddig3 && udig3 > udig1 && udig3 > ddig1 && udig3 > udig2 && udig3 > ddig2 {
							fmt.Println("Se determina que el dígito mayor: " + udig3 + " se encuentra en número num3: " + num3)
						} else {
							if ddig3 > udig3 && ddig3 > udig1 && ddig3 > ddig1 && ddig3 > udig2 && ddig3 > ddig2 {
								fmt.Println("Se determina que el dígito mayor: " + ddig3 + " se encuentra en número num3: " + num3)
							}
						}
					}
				}
			}
		}
	} else {
		fmt.Println("Solo uno o ninguno de los 3 números es de 2 dígitos ")
	}
}
