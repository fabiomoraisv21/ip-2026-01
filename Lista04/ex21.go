package main

import "fmt"

var nn, codigo int

func main() {
	vetor := []int{}
	fmt.Scan(&codigo)
	for i := 0; i < 10; i++ {
		fmt.Scan(&nn)
		vetor = append(vetor, nn)
	}
	if codigo == 0 {
		return
	} else if codigo == 1 {
		fmt.Println(vetor)
	} else if codigo == 2 {
		vetorInverso := []int{}
		for i := 9; i >= 0; i-- {
			vetorInverso = append(vetorInverso, vetor[i])
		}
		fmt.Println(vetorInverso)
	}
}
