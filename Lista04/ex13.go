package main

import "fmt"

func main() {
	var ids []int
	var meses []int

	fmt.Println("Digite o ID e os Meses (0 0 para encerrar):")

	for i := 0; i < 100; i++ {
		var id, mes int
		fmt.Scan(&id, &mes)

		// Condição de parada
		if id == 0 && mes == 0 {
			break
		}

		ids = append(ids, id)
		meses = append(meses, mes)
	}

	tamanho := len(meses)

	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho-1-i; j++ {
			if meses[j] > meses[j+1] {
				tempMes := meses[j]
				meses[j] = meses[j+1]
				meses[j+1] = tempMes

				tempId := ids[j]
				ids[j] = ids[j+1]
				ids[j+1] = tempId
			}
		}
	}

	fmt.Println("\nOs três empregados mais recentes:")

	limite := 3
	if len(ids) < 3 {
		limite = len(ids)
	}

	for i := 0; i < limite; i++ {
		fmt.Printf("%d° mais recente: ID %d com %d meses\n", i+1, ids[i], meses[i])
	}
}
