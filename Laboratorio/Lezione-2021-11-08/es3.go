/*
TESTO DEL PROBLEMA
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	fmt.Println("Inserisci un numero:")
	fmt.Scan(&x)
	radiceQuadrata := math.Sqrt(x)

	if x == radiceQuadrata*radiceQuadrata {
		fmt.Println(x, "uguale a", radiceQuadrata, "*", radiceQuadrata)
	} else {
		fmt.Println(x, "diverso da", radiceQuadrata, "*", radiceQuadrata)
	}
}
