package main

import "fmt"

var numeroo int

func main() {
	fmt.Println("Digite um número:")
	fmt.Scan(&numeroo)
	if numeroo < 0 {
		fmt.Println("Esse número é inválido")
	} else {
		if numeroo == 0 {
			fmt.Println("O fatorial de 0 é 1")
		} else {
			for i := (numeroo - 1); i > 0; i-- {
				numeroo *= i
			}
			fmt.Println("O fatorial desse número é", numeroo)
		}
	}
}
