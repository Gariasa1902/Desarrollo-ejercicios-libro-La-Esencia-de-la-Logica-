package main

import "fmt"

func main() {
	// Leer un número entero y determinar cuántos dígitos tiene.
	var num int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num)

	var cont int16 = 0
	for num > 0 {
		num = num / 10
		cont++
	}
	fmt.Println("El número ingresado tiene ", cont, " dígitos ")
}
