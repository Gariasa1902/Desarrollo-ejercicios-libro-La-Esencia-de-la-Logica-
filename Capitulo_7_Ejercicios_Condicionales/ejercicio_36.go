package main

import "fmt"

func main() {
	//36. Leer un número entero de 4 dígitos y determinar si tiene mas dígitos pares o impares.
	var num int32
	fmt.Println("Ingrese un número entero de 4 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	var par int32 = 0
	var impar int32 = 0
	if num >= 1000 && num <= 9999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100
		var umdig int32 = (num % 10000) / 1000
		if (udig/2)*2 == udig {
			par++
		} else {
			impar = impar + 1
		}
		if (ddig/2)*2 == ddig {
			par = par + 1
		} else {
			impar = impar + 1
		}
		if (cdig/2)*2 == cdig {
			par = par + 1
		} else {
			impar = impar + 1
		}
		if (umdig/2)*2 == umdig {
			par = par + 1
		} else {
			impar = impar + 1
		}
		fmt.Println("Se determina que el número ingresado tiene ", par, " dígitos pares ")
		fmt.Println("Se determina que el número ingresado tiene ", impar, " dígitos pares ")
		if par > impar {
			fmt.Println("Se determina que el número ingresado tiene mas dígitos pares ")
		} else {
			if impar > par {
				fmt.Println("Se determina que el número ingresado tiene mas dígitos impares ")
			} else {
				if par == impar {
					fmt.Println("Se determina que el número ingresado tiene la misma cantidad de dígitos pares e impares ")
				}
			}
		}
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}

}
