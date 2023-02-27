package main

import (
    "fmt"
    "time"
)

func main() {
    // Get the current time
    t := time.Now()
  
    // Get the day of the week
    dayOfWeek := t.Weekday()

    // Print the day of the week
    fmt.Printf("Today is %s\n", dayOfWeek)
	fmt.Printf("the time is %s", t.Format("15:04 PM"))
}
