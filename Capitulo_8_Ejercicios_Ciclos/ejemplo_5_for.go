package main

import "fmt"

func main() {
	/* 3. Leer dos números enteros y determinar cual de los dos tiene mas
	dígitos*/
	var num1, num2 int32
	fmt.Println("Ingrese el primer número entero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo número entero: ")
	fmt.Scan(&num2)
	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1) // Convertir a positivo si es negativo
		num2 = num2 * (-1) // Convertir a positivo si es negativo
	}
	var cont1 int32 = 0
	var i int32
	for i = 0; i != num1; cont1++ {
		num1 = num1 / 10 // Elimina el último dígito
	}
	fmt.Println("Se determina que el primer número ingresado tiene ", cont1, " dígitos ")

	var cont2 int32 = 0
	for i = 0; i != num2; cont2++ {
		num2 = num2 / 10 // Elimina el último dígito
	}
}
