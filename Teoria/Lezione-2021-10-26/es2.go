/*
ENUMERARE I FALSI NEGATIVI DELLA PROVA DEL 9
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			for r := 1; r < n*n; r++ { //Provo tutti i risultati, per vedere quanti tra quelli sbagliati vengono verificati dalla prova del 9
				if r != a*b && r%9 == (a*b)%9 {
					fmt.Println(a, b, r)
				}
			}
		}
	}
}
