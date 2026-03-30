// 18. PA
package main

import "fmt"

func main() {
	var a1, r, n int

	fmt.Println("Digite o primeiro termo:")
	fmt.Scan(&a1)

	fmt.Println("Digite a razao:")
	fmt.Scan(&r)

	fmt.Println("Digite a quantidade de termos:")
	fmt.Scan(&n)

	soma := 0

	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}

	fmt.Println("SOMA =", soma)
}
