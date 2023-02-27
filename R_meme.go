package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    topTexts := []string{"One does not simply", "But that's none of my business", "Not sure if", "Why not both?", "I don't always"}
    bottomTexts := []string{"but when I do, I", "grumpy cat", "skeptical third world kid", "brace yourself", "get shit done"}

    topIndex := rand.Intn(len(topTexts))
    bottomIndex := rand.Intn(len(bottomTexts))

    meme := fmt.Sprintf("%s\n%s", topTexts[topIndex], bottomTexts[bottomIndex])
    fmt.Println(meme)
}
