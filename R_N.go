package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Set the seed value for the random number generator
    rand.Seed(time.Now().UnixNano())

    // Generate and print 10 random integers between 0 and 99
    for i := 0; i < 10; i++ {
        fmt.Println(rand.Intn(100))
    }
}