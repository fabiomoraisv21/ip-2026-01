package main

import "fmt"

var numero int

func main() {
	for {
		fmt.Println("Digite um número:")
		fmt.Scan(&numero)
		if numero <= 0 {
			fmt.Println("Esse número é inválido")
			break
		}
		for i := 1; i <= numero; i++ {
			if numero/i == i {
				fmt.Println("Esse número é uma raiz perfeita")
				break
			} else {
				continue
			}
		}
	}
}
