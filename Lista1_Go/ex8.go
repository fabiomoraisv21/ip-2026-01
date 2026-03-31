// 8. Lata
package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h float64

	fmt.Println("Digite o raio da lata (em metros):")
	fmt.Scan(&r)

	fmt.Println("Digite a altura da lata (em metros):")
	fmt.Scan(&h)

	area := 2*math.Pi*r*r + 2*math.Pi*r*h
	custo := area * 100

	fmt.Println("O VALOR DO CUSTO É =", custo)
}
