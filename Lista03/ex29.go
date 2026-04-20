package main

import "fmt"

var n, somaa int

func main() {
	fmt.Println("Digite o valor de N:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		somaa += i
	}
	fmt.Printf("A soma de 1 até N é %v\n", somaa)
}
