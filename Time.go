package main

import (
	"fmt"
	"time"
)

func whatdayisit() {
	fmt.Println("Saturday")
}

func main() {
	go whatdayisit()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}