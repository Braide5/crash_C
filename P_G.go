package main

import (
	"log"
	"math/rand"
	"fmt"

	
)

func main() {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower casae letters, disallowing repeat characyers.
	res, err := math/rand.int
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(res)
	fmt.Println(rand.Intn(200))
}