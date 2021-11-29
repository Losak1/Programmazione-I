/*
TESTO DEL PROBLEMA
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var testo string
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return testo[:len(testo)-1]
}

func SeparaParole(s string) []string {
	parole := []string{}
	for _, line := range strings.Split(s, "\n") {
		for _, word := range strings.Split(line, " ") {
			if word != "" {
				parole = append(parole, word)
			}
		}
	}
	return parole
}

func ContaRipetizioni(sl []string) map[string]int {
	rip := make(map[string]int)
	for _, w := range sl {
		if _, ok := rip[w]; !ok {
			rip[w] = 1
		} else {
			rip[w]++
		}
	}
	return rip
}

func main() {

	testo := LeggiTesto()
	parole := SeparaParole(testo)
	ripetizioni := ContaRipetizioni(parole)

	fmt.Print("Numero di parole distinte: ", len(ripetizioni), "\n")
	for k, v := range ripetizioni {
		fmt.Print(k, ": ", v, "\n")
	}

}
