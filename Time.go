package main

import (
    "fmt"
    "time"
)

func main() {
    // Get the current time
    now := time.Now()

    // Get the day of the week
    dayOfWeek := now.Weekday()

    // Print the day of the week
    fmt.Printf("Today is %s\n", dayOfWeek)
	fmt.Printf("the time is %d", now)
}
