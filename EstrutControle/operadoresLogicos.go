package main

import "fmt"

func main() {
	// se X igual a 2 ou 3, print "sim" se nao print "nao"
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("Sim, X e 2 ou 3")
	} else {
		fmt.Println(("X nao e nem 2 e nem 3"))
	}

	// se X for multiplo de 4 print "sim" se nao print "nao"
	a := 4

	if a%2 == 0 && a%3 == 0 {
		fmt.Println("A e multiplo de 2 nem 3")
	} else {
		fmt.Println("A nao multipo de 2 nem 3")
	}

}

// && = e
// || = ou
// ! = negação
