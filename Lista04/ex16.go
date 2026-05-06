package main

import "fmt"

var idades int

func moda(array []int, n int) int {
	count := 0
	for _, v := range array {
		if v == n {
			count++
		}
	}
	return count
}

func main() {
	modaa := 0
	idadeModa := 0
	vetorIdades := []int{}
	for i := 0; i < 50; i++ {
		fmt.Scan(&idades)
		vetorIdades = append(vetorIdades, idades)
	}
	for i := 0; i < 50; i++ {
		if moda(vetorIdades, vetorIdades[i]) > modaa {
			idadeModa = vetorIdades[i]
			modaa = moda(vetorIdades, vetorIdades[i])
		}
	}
	fmt.Println(idadeModa)
}
