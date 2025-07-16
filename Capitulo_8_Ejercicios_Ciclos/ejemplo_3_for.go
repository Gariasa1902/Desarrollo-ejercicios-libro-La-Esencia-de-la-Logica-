package main

import "fmt"

func main() {
	// 1. Leer un número entero y determinar cuántos dígitos tiene
	var num int32
	fmt.Println("Ingrese un número entero: ")
	fmt.Scan(&num)
	if num < 0 {
		num = num * (-1) // Convertir a positivo si es negativo
	}
	var cont int32 = 0
	var i int32
	for i = 0; num != i; cont++ {
		num = num / 10 // Eliminar el último dígito
	}
	fmt.Println("El número ingresado tiene ", cont, " dígitos ")
}
