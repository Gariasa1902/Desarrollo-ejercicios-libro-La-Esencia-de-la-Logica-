package main

import "fmt"

func main() {
	// 23. Leer un número entero de tres dígitos y determinar cuántos dígitos primos tiene.
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
		if udig == 2 || udig == 3 || udig == 5 || udig == 7 {
			cont++
		}
		if ddig == 2 || ddig == 3 || ddig == 5 || ddig == 7 {
			cont++
		}
		if cdig == 2 || cdig == 3 || cdig == 5 || cdig == 7 {
			cont++
		}
		fmt.Println("Se determina que el número ingresado tiene ", cont, " números primos")
	} else {
		fmt.Println("El número ingresado no es de 3 dígitos ")
	}
}
