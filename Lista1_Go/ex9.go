// 9. Delta
package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Println("Digite o valor de A:")
	fmt.Scan(&a)

	fmt.Println("Digite o valor de B:")
	fmt.Scan(&b)

	fmt.Println("Digite o valor de C:")
	fmt.Scan(&c)

	d := b*b - 4*a*c

	fmt.Println("O VALOR DE DELTA É =", d)
}
