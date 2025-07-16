package main

import "fmt"

func main() {
	/* 47. Leer dos números enteros y si la diferencia entre los dos números es par mostrar en pantalla la
	suma de los dígitos de los números, si dicha diferencia es un número primo menor que 10
	entonces mostrar en pantalla el producto de los dos números y si la diferencia entre ellos
	terminar en 4 mostrar en pantalla todos los dígitos por separado*/
	var num1 int32
	fmt.Println("Ingrese un número entero ")
	fmt.Scan(&num1)
	var num2 int32
	fmt.Println("Ingrese otro número entero ")
	fmt.Scan(&num2)

	if num1 < 0 && num2 < 0 {
		num1 = num1 * (-1)
		num2 = num2 * (-1)
	}
	if num1 > num2 {
		var dif int32 = num1 - num2

		var suma1 int32 = 0
		var suma2 int32 = 0
		if (dif/2)*2 == dif {
			for num1 != 0 {
				var udig1 int32 = num1 % 10
				suma1 = suma1 + udig1
				num1 = num1 / 10
			}
			for num2 != 0 {
				var udig2 int32 = num2 % 10
				suma2 = suma2 + udig2
				num2 = num2 / 10
			}
			fmt.Println("Se determina que la diferencia de los 2 números ingresados es un numero par: ", dif, " y la suma de todos los dígitos es: ", suma1+suma2)
		}

		var i int32 = 1 //CONTROLA EL CICLO WHILE
		var cont int32 = 0
		for i <= dif {
			if dif%i == 0 {
				cont = cont + 1
			}
			i = i + 1
		}
		if cont == 2 && cont < 10 {
			fmt.Println("Se determina que la diferencia de los 2 números ingresados es un numero primo, menor que 10 y el producto de todos los dígitos es: ", num1*num2)
		}

		var udif int32 = dif % 10
		var cont2 int32 = 0
		//var cont3 int32 = 0
		if udif == 4 {
			for dif != 0 {
				dif = dif / 10
				cont2 = cont2 + 1
			}
			fmt.Println(cont2)
			fmt.Println("resultado ", dif)
		}
	} else {
		if num2 > num1 {
			var dif int32 = num2 - num1

			var suma1 int32 = 0
			var suma2 int32 = 0
			if (dif/2)*2 == dif {
				for num1 != 0 {
					var udig1 int32 = num1 % 10
					suma1 = suma1 + udig1
					num1 = num1 / 10
				}
				for num2 != 0 {
					var udig2 int32 = num2 % 10
					suma2 = suma2 + udig2
					num2 = num2 / 10
				}
				fmt.Println("Se determina que la diferencia de los 2 números ingresados es un numero par: ", dif, " y la suma de todos los dígitos es: ", suma1+suma2)
			}

			var i int32 = 1 //CONTROLA EL CICLO WHILE
			var cont int32 = 0
			for i <= dif {
				if dif%i == 0 {
					cont = cont + 1
				}
				i = i + 1
			}
			if cont == 2 && cont < 10 {
				fmt.Println("Se determina que la diferencia de los 2 números ingresados es un numero primo, menor que 10 y el producto de todos los dígitos es: ", num1*num2)
			}

			var udif int32 = dif % 10
			var cont2 int32 = 0
			//var cont3 int32 = 0
			if udif == 4 {
				for dif != 0 {
					dif = dif / 10
					cont2 = cont2 + 1
				}
				fmt.Println(cont2)
				fmt.Println("resultado ", dif)
			}
		}
	}
}
