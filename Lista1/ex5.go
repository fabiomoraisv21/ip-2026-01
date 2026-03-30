// 5. Água
package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string
	var valor float64

	fmt.Println("Digite o identificador da conta, o consumo e a categoria (R, C ou I):")
	fmt.Scan(&conta, &consumo, &tipo)

	if tipo == "R" {
		valor = 5 + consumo*0.05
	} else if tipo == "C" {
		valor = 500 + consumo*0.25
	} else {
		valor = 800 + consumo*0.04
	}

	fmt.Println("CONTA =", conta)
	fmt.Println("VALOR DA CONTA =", valor)
}
