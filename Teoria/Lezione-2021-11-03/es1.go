/*
ESEMPI FUNZIONAMENTO PUNTATORI
*/

package main

import "fmt"

func main() {

	var x int = 5
	var p *int = &x
	var q **int = &p
	fmt.Println("x", x)
	fmt.Println("p", p)
	fmt.Println("q", q)
	fmt.Println("*p", *p)
	fmt.Println("*q", *q)
	fmt.Println("&x", &x)
	fmt.Println("&p", &p)
	fmt.Println("&q", &q)
}
