package main

import "fmt"

var valor int

func main() {
	vetor := [10]int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&valor)

		j := i - 1
		for j >= 0 && vetor[j] > valor {
			vetor[j+1] = vetor[j]
			j--
		}
		vetor[j+1] = valor
	}
	fmt.Println(vetor)
}
