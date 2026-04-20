package main

import "fmt"

var termo int

func main() {
	sequencia := []int{}
	fmt.Println("Quantos termos quer na sequência?")
	fmt.Scan(&termo)
	for i := 1; i <= termo; i++ {
		sequencia = append(sequencia, i*i)
	}
	fmt.Println(sequencia)
}
