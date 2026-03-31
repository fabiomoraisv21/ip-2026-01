// 20. Tempo
package main

import "fmt"

func main() {
	var h, m, s int

	fmt.Println("Digite as horas:")
	fmt.Scan(&h)

	fmt.Println("Digite os minutos:")
	fmt.Scan(&m)

	fmt.Println("Digite os segundos:")
	fmt.Scan(&s)

	total := h*3600 + m*60 + s

	fmt.Println("O TEMPO EM SEGUNDOS É =", total)
}
