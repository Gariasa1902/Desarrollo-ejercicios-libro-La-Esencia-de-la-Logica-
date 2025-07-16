package main

import "fmt"

func main() {
	// Escribir los números impares comprendidos entre 1 y 20.
	for num := 1; num <= 20; num++ {
		if num%2 != 0 {
			fmt.Println(num)
		}
	}
}
