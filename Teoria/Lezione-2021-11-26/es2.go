/*
TESTO DEL PROBLEMA
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	s := []float64{0.2, 1.2, 2, 3, 4, 5}
	rand.Shuffle(len(s), func(i, j int) { //andrebbe cambiato il seme
		s[i], s[j] = s[j], s[i]
	})

	fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}
