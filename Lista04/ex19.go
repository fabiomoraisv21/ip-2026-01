package main

import "fmt"

var elemento, divisor int

func main() {
	vetor := []int{}
	divisores := []int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&elemento)
		vetor = append(vetor, elemento)
	}
	for i := 0; i < 5; i++ {
		fmt.Scan(&divisor)
		divisores = append(divisores, divisor)
	}
	for j := 0; j < 10; j++ {
		fmt.Println("Número", vetor[j])
		for i, v := range divisores {
			if vetor[j]%v == 0 {
				fmt.Printf("\tDivisível por %v na posição %v\n", v, i)
			} else {
				continue
			}
		}
	}
}
