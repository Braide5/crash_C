package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Printf("The time is %s\n", t.Format("15:04 PM"))
}
