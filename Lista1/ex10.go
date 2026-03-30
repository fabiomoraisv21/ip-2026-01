// 10. Determinante
package main

import "fmt"

func main() {
	var a, b, c, d float64

	fmt.Println("Digite o valor de a:")
	fmt.Scan(&a)

	fmt.Println("Digite o valor de b:")
	fmt.Scan(&b)

	fmt.Println("Digite o valor de c:")
	fmt.Scan(&c)

	fmt.Println("Digite o valor de d:")
	fmt.Scan(&d)

	det := a*d - b*c

	fmt.Println("O VALOR DO DETERMINANTE É =", det)
}
