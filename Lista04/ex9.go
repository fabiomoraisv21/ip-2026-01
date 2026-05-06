
package main

import "fmt"

func main() {

	arr := make([]float64, 15)

	soma := 0.0

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a %d° altura: ", i+1)
		fmt.Scan(&arr[i])
		soma += arr[i]
	}

	media := soma / 15
	var arr2 []float64

	for _, altura := range arr {
		if altura > media {
			arr2 = append(arr2, altura)
		}
	}

	fmt.Print(arr2)
}
