package main

import "fmt"

func main() {
	A := 1
	for A <= 10 {
		fmt.Println(A)
		A = A + 1
	}

	fmt.Println("até 20")

	//segunda forma de fazer o for
	//i++  pegar o valor da variavel e somar +1

	for i := 11; i <= 20; i++ {
		fmt.Println(i)
	}

}
