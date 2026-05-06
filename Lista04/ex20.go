package main

import "fmt"

var valorDado int

func contaa(array []int, n int) int {
	count := 0
	for _, v := range array {
		if v == n {
			count++
		}
	}
	return count
}

func main() {
	vetor := []int{}
	for i := 0; i < 20; i++ {
		fmt.Scan(&valorDado)
		vetor = append(vetor, valorDado)
	}
	for i := 0; i < 20; i++ {
		fmt.Printf("Número %v apareceu %v vezes\n", vetor[i], contaa(vetor, vetor[i]))
	}
}
