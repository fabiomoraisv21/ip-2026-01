
package main

import (
	"fmt"
	"math"
)

func main() {

	var arr [15]int
	var raiz [15]float64

	for i := 0; i < 15; i++{
		fmt.Scan(&arr[i])
		if arr [i] >= 0{
			raiz[i] = math.Sqrt(float64(arr[i]))
		} else {
			raiz[i] = -1
		}
	}
	fmt.Print(raiz)
}