package main

import (
  "fmt"
  "time"
)

func main() {
  ch1 := make(chan string)
  ch2 := make(chan string)

  // Goroutine sending to ch1
  go func() {
    time.Sleep(1 * time.Second)
    ch1 <- "Message from channel 1"
  }()

  // Goroutine sending to ch2
  go func() {
    time.Sleep(2 * time.Second)
    ch2 <- "Message from channel 2"
  }()

  // Select waits on multiple channels
  for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
      fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
      fmt.Println("Received:", msg2)
    }
  }

  fmt.Println("Done!")
}

