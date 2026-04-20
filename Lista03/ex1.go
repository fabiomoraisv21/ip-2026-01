package main

import "fmt"

var base, expoente, resultado int

func main() {
	fmt.Println("Digite a base:")
	fmt.Scan(&base)
	fmt.Println("Digite o expoente")
	fmt.Scan(&expoente)
	resultado = base
	for i := 1; i < expoente; i++ {
		resultado *= base
	}
	fmt.Printf("O resultado é %v\n", resultado)
}
