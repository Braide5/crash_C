package main

import (
	"fmt"
	"math/rand"
	"time"
)

var symbols = []rune{'@', '#', '$'}

func main() {
	rand.Seed(time.Now().UnixNano())
	password := make([]rune, 6)
	
	for i := range password {
		switch rand.Intn(4) {
		case 0:
			password[i] = rune(rand.Intn(10) + '0') // digit
		case 1:
			password[i] = rune(rand.Intn(26) + 'A') // uppercase letter
		case 2:
			password[i] = rune(rand.Intn(26) + 'a') // lowercase letter
		case 3:
			password[i] = symbols[rand.Intn(len(symbols))] // symbol
		}
	}
	fmt.Printf("Random password: %s\n", string(password))
}
