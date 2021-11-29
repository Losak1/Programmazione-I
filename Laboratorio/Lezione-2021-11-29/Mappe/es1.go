/*
TESTO DEL PROBLEMA
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var testo string
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return testo[:len(testo)-1]
}

func Occorrenze(s string) map[rune]int {
	occ := make(map[rune]int)
	for _, l := range s {
		if unicode.IsLetter(l) {
			if _, ok := occ[l]; !ok {
				occ[l] = 1
			} else {
				occ[l]++
			}
		}
	}
	return occ
}

func Asterischi(x int) string {
	s := ""
	for i := 0; i < x; i++ {
		s += "*"
	}
	return s
}

func StampaIstogramma(occorrenze map[rune]int) {
	lettere := []rune{}
	for k := range occorrenze {
		lettere = append(lettere, k)
	}

	sort.Slice(lettere, func(i, j int) bool {
		return lettere[i] < lettere[j]
	})

	for _, k := range lettere {
		fmt.Print(string(k), ": ", Asterischi(occorrenze[k]), "\n")
	}

}

func main() {

	testo := LeggiTesto()
	occ := Occorrenze(testo)
	StampaIstogramma(occ)

}
