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

func StatisticheParole(s string) (int, int, int) {
	var linee []string = strings.Split(s, "\n")
	for i := range linee {
		strings.Split(linee[i], " ")
	}
	return 0, 0, 0
}

func main() {

	testo := LeggiTesto()
	numeroLinee, numeroParole, lunghezzaTotale := StatisticheParole(testo)
	fmt.Println(numeroLinee)
	fmt.Println(numeroParole)
	fmt.Println(lunghezzaTotale)

}
