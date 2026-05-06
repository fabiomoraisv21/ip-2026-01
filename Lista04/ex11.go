package main

import "fmt"

var numm, somatorio int

func main() {
	vetor := []int{}
	for i:=0; i < 100; i++ {
		fmt.Scan(&numm)
		vetor = append(vetor, numm)
	}
	for i:=0; i <= 49; i++ {
		resultado := (vetor[i]-vetor[99-i]) * (vetor[i]-vetor[99-i]) * (vetor[i]-vetor[99-i])
		somatorio += resultado
	}
	fmt.Println(somatorio)
}