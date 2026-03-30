// 14. Pirâmide
package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64

	fmt.Println("Digite a altura da pirâmide (m):")
	fmt.Scan(&h)

	fmt.Println("Digite o lado do hexágono (m):")
	fmt.Scan(&a)

	ab := (3 * a * a * math.Sqrt(3)) / 2
	volume := (ab * h) / 3

	fmt.Println("VOLUME =", volume, "METROS CÚBICOS")
}
