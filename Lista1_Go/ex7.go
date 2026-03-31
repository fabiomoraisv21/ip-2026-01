// 7. Métrico
package main

import "fmt"

func main() {
	var f, polegada float64

	fmt.Println("Escreva a temperatura em Fahrenheit:")
	fmt.Scan(&f)

	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scan(&polegada)

	c := (5 * (f - 32)) / 9
	mm := polegada * 25.4

	fmt.Println("O VALOR EM CELSIUS =", c)
	fmt.Println("A QUANTIDADE DE CHUVA É =", mm)
}
