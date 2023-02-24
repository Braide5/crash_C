package main

import (
	"log"
	"math/rand"
	"fmt"

	
)

func main() {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower casae letters, disallowing repeat characyers.
	res := rand.Intn(2000000000)
	
	log.Printf("%d", res)
	fmt.Println(res)
	fmt.Println(rand.Intn(2000000000))
}