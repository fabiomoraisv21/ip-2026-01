package main

import "fmt"

var pos1, pos2, nMultiplo, multiplo1, multiplo2 int

func main() {
	fmt.Println("Digite o valor do primeiro número:")
	fmt.Scan(&pos1)
	fmt.Println("Digite o valor do segundo número:")
	fmt.Scan(&pos2)
	mmc := pos1
	if pos2 > pos1 {
		mmc = pos2
	}

	for {
		if mmc%pos1 == 0 && mmc%pos2 == 0 {
			break
		}
		mmc++
	}
	fmt.Printf("O MMC entre os números é %v\n", mmc)

}
