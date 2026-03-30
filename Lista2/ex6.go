// Exercicio 6

package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("Digite um número A, e um número B:")
	fmt.Scan(&n1, &n2)

	if n1%n2 == 0 {
		fmt.Printf("O número A:%v é divisível por B:%v", n1, n2)
	} else {
		fmt.Printf("O número A:%v não é divisível por B:%v", n1, n2)
	}
}
