// 6. Conversão Temperatura
package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite a quantidade de temperaturas:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var f float64

		fmt.Println("Digite a", i, "ª temperatura em Fahrenheit:")
		fmt.Scan(&f)

		c := (5 * (f - 32)) / 9

		fmt.Println(f, "FAHRENHEIT EQUIVALE A", c, "CELSIUS")
	}
}
