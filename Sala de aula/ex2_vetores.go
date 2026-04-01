package main

import "fmt"

func main() {

	var numeros [5]int
	var soma int

	for i := 0; i < 5; i++ {

		fmt.Printf("Digite o %d° numero inteiro:", i+1)
		fmt.Scan(&numeros[i])
		soma += numeros[i]
	}

	fmt.Printf("A soma total dos números inteiros é: %d", soma)

}
