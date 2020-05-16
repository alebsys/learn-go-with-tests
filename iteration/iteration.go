package main

import (
	"fmt"
	"os"
)

const repeatCount = 5

// Repeated func
func Repeated(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	var number int
	fmt.Fscan(os.Stdin, &number)
	character := "a"
	fmt.Println(Repeated(character, number))
}
