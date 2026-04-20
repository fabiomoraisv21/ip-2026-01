package main

import "fmt"

var soma1ao20 int

func main() {
	for i := 1; i <= 20; i++ {
		soma1ao20 += i
	}
	fmt.Printf("A soma dos números entre 1 e 20 é %v\n", soma1ao20)
}
