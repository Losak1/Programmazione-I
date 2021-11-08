/*
CONTA LE OCCORRENZE DI UNA STRINGA IN UN'ALTRA
*/

package main

import "fmt"

func cerca(s, p string) int {
	c := 0
	for i := 0; i < len(s)-len(p)+1; i++ {
		j := 0
		for ; j < len(p); j++ {
			if p[j] != s[i+j] {
				break
			}
		}
		if j == len(p) { //se vera, è uscito alla fine, quindi è un'occorrenza
			c++
		}
	}
	return c
}

func main() {

	var s, p string
	fmt.Print("Stringa: ")
	fmt.Scan(&s)
	fmt.Print("Pattern: ")
	fmt.Scan(&p)

	fmt.Println(cerca(s, p))
}
