package main

import "fmt"

func main() {
	// 24. Leer un número entero de tres dígitos y determinar cuántos dígitos pares tiene.
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

		var cont int32 = 0
		if (udig/2)*2 == udig {
			cont++
		}
		if (ddig/2)*2 == ddig {
			cont++
		}
		if (cdig/2)*2 == cdig {
			cont++
		}
		fmt.Println("Se determina que el número ingresado tiene ", cont, " números pares")
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
