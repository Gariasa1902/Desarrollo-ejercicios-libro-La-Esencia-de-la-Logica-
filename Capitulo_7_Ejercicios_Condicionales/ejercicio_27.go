package main
import "fmt"
func main(){
	// 27. Leer un número entero de cuatro dígitos y determinar cuántos dígitos pares tiene.
	var num int32
	fmt.Println("Ingrese un número entero de 4 dígitos ")
	fmt.Scan(&num)

	if num < 0 {
		num = num * (-1)
	}

	if num >= 1000 && num <= 9999 {
		var udig int32 = num % 10
		var ddig int32 = (num % 100) / 10
		var cdig int32 = (num % 1000) / 100
		var ucdig int32 = (num % 10000) / 1000

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
		if (ucdig/2)*2 == ucdig
		fmt.Println("Se determina que el número ingresado tiene ", cont, " números pares")
	} else {
		fmt.Println("El número ingresado no es de 4 dígitos ")
	}
}