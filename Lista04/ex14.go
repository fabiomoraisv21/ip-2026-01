package main

import "fmt"

var nVetor1, nVetor2 int

func main() {
	vetor1 := []int{}
	vetor2 := []int{}
	juncao := []int{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&nVetor1)
		vetor1 = append(vetor1, nVetor1)
	}
	for i := 0; i < 10; i++ {
		fmt.Scan(&nVetor2)
		vetor2 = append(vetor2, nVetor2)
	}
	for i := 0; i < 10; i++ {
		juncao = append(juncao, vetor1[i])
		juncao = append(juncao, vetor2[i])
	}
	fmt.Println(juncao)
}
