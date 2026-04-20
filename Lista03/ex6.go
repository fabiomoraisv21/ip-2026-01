package main

import "fmt"

var num int

func main() {
	fmt.Println("Digite um número inteiro e positivo:")
	fmt.Scan(&num)
	triangular := false
	for i := 1; i <= num; i++ {
		if i*(i+1)*(i+2) == num {
			triangular = true
			break
		}
	}
	if triangular == true {
		fmt.Println("Esse número é triangular")
	} else {
		fmt.Println("Esse número não é triangular")
	}
}
