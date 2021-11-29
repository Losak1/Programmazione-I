/*
TESTO DEL PROBLEMA
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var s1 []int

	for i := 1; i < len(os.Args); i++ {
		x, err := strconv.Atoi(os.Args[i])
		if err == nil {
			s1 = append(s1, x)
		}
	}

	if len(s1) != 0 {
		min := Minimo(s1)
		fmt.Println("Minimo:", min)
		max := Massimo(s1)
		fmt.Println("Massimo:", max)
		avg := Media(s1)
		fmt.Println("Valore medio:", avg)
	}
}

func Minimo(sl []int) int {
	M := math.MaxInt
	for _, x := range sl {
		if x < M {
			M = x
		}
	}
	return M
}

func Massimo(sl []int) int {
	M := math.MinInt
	for _, x := range sl {
		if x > M {
			M = x
		}
	}
	return M
}

func Media(sl []int) float64 {
	somma := 0
	for _, x := range sl {
		somma += x
	}
	media := float64(somma) / float64(len(sl))
	return media
}
