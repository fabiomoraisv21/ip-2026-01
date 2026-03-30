// 13. Conceito
package main

import "fmt"

func main() {
	var nota float64

	fmt.Println("Digite a nota (0 a 10):")
	fmt.Scan(&nota)

	var conceito string

	if nota >= 9 {
		conceito = "A"
	} else if nota >= 7.5 {
		conceito = "B"
	} else if nota >= 6 {
		conceito = "C"
	} else {
		conceito = "D"
	}

	fmt.Println("NOTA =", nota, "| CONCEITO =", conceito)
}
