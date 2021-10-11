package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	var r int

	r = a - b

	fmt.Println(r)

}

/*
Supponendo che l'input fornito da standard input sia

5 6

cosa dovrebbe produrre in output il seguente programma?

-1 cioè la differenza tra il primo numero e il secondo

*/
