package main

import "fmt"

const lowerLimit = 1
const upperLimit = 100

func main() {
	for i := lowerLimit; i < upperLimit; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}
