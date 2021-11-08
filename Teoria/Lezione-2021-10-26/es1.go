/*
INDIVIDUARE TUTTE LE TERNE PITAGORICHE CON LATI SOTTO UN CERTO N
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			r := math.Sqrt(float64(a*a + b*b))
			if r == float64(int(r)) {
				fmt.Println(a, b, r)
			}
		}
	}
}
