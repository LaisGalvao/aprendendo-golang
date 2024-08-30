package main

import (
	"fmt"
	"strconv"
)

func sayTheName(s string) {
	fmt.Println("Olá", s)
}

func convertStringToNumber(s string) (n int) {
	i, _ := strconv.Atoi(s)
	n = i
	return
}

func main() {
	n := convertStringToNumber("123")
	fmt.Println("O número convertido é:", n)
	sayTheName("Laís")
}
