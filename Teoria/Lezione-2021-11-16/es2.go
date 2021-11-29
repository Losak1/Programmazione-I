/*
FUNZIONE CHE CREI UNA MAPPA CON LE APPARENZE DELLE RUNE IN UNA STRINGA
*/

package main

import "fmt"

func stat(s string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, r := range s {
		if v, ok := m[r]; ok {
			m[r] = v + 1 //v è m[r]
		} else {
			m[r] = 1
		}
	}
	return m
}

func main() {

	var s string

	fmt.Scan(&s)

	for k, v := range stat(s) {
		fmt.Println(string(k), v)
	}

}
