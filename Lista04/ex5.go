package main

import "fmt"

func main () {

	arr := make([]int, 10)

	for i := 0; i < 10; i++{
		fmt.Scan(&arr[i])
	}

	seilavey := arr[0]
	hanna := 0

	for j := 1; j < 10; j++{
		if arr[j] < seilavey{
			seilavey = arr[j]
			hanna = j
		}
	}
	fmt.Print("O menor valor é: ",seilavey," e sua posição é: ",hanna + 1,"°")
}