package main

import "fmt"

const lowerLimit = 1
const upperLimit = 100

func main() {
	for i := lowerLimit; i < upperLimit; i++ {
		// O exercício não fala em casos que o número é
		// multiplo de 3 e 5, como 15. Nesses casos apenas 
		// a primeira condicional vai ser atingida.
		if i % 3 == 0 {
			fmt.Println(i, "- Pin")
		} else if i % 5 == 0 {
			fmt.Println(i, "- Pan")
		}
	}
}
