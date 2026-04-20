package main

import "fmt"

var n1, n2, m int

func main() {
	fmt.Println("Digite o valor de N1:")
	fmt.Scan(&n1)
	fmt.Println("Digite o valor de N2:")
	fmt.Scan(&n2)
	for i := 0; i < n2; i++ {
		m += n1
	}
	fmt.Printf("A multiplicação entre N1 e N2 é %v\n", m)
}
