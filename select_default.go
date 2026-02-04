package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)

  // Send after delay
  go func() {
    time.Sleep(2 * time.Second)
    ch <- "Hello!"
  }()

  // Non-blocking receive with default
  for i := 0; i < 5; i++ {
    select {
    case msg := <-ch:
      fmt.Println("Received:", msg)
      return
    default:
      fmt.Println("No message yet, waiting...")
      time.Sleep(500 * time.Millisecond)
    }
  }

  // Non-blocking send example
  messages := make(chan string, 1)
  messages <- "buffered"

  select {
  case messages <- "overflow":
    fmt.Println("Sent overflow message")
  default:
    fmt.Println("Channel full, message dropped")
  }
}

