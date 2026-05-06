package main

import "fmt"

var inteiroo int

func main() {
	vetorde30 := []int{}
	novoVetor := []int{}
	for i := 0; i < 30; i++ {
		fmt.Scan(&inteiroo)
		vetorde30 = append(vetorde30, inteiroo)
	}
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			novoVetor = append(novoVetor, vetorde30[i]*2)
		} else {
			novoVetor = append(novoVetor, vetorde30[i]*3)
		}
	}
	fmt.Println(novoVetor)
}
