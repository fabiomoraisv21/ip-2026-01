package main

import "fmt"

func main() {
	arr := make([]int, 10)

	for i := range 10 {
		fmt.Scan(&arr[i])
	}

	encontrou := false

	for j := 0; j < 10; j++ {
		if arr[j] > 50 {
			fmt.Printf("O número %d° é: %d \n", j+1, arr[j])
			encontrou = true
		}
	}
	if encontrou == false {
		fmt.Print("Não há números maiores que 50")
	}
}
