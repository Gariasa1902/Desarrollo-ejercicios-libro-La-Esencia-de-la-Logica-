package main

import "fmt"

func main() {
	// Escribir los números impares comprendidos entre 1 y 20.
	var num int32 = 1
	for num <= 20 {
		if num%2 != 0 {
			fmt.Println(num)
		}
		num++
	}
}
