/*
ESEMPI FUNZIONAMENTO PUNTATORI
*/

package main

import "fmt"

func main() {

	var x int = 5
	fmt.Println(x)
	var p *int
	p = &x
	*p *= 5
	fmt.Println(x)
	var q *int = new(int)
	*q = 15
	println(*q)
}
