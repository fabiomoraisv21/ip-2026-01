package main

import "fmt"

var nota int

func conta(slice []int, value int) int {
	count := 0
	for _, v := range(slice) {
		if v == value {
			count++
		}
	}
	return count
}

func main() {
	vetor := []int{}
	for i:=0; i < 15; i++ {
		fmt.Scan(&nota)
		vetor = append(vetor, nota)
	}
	for _, v := range(vetor) {
		freq := conta(vetor, v)
		freqAbs := float64(conta(vetor, v))/15
		fmt.Println(v, freq, freqAbs)
	}
}