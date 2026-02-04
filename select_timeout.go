package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)

  // Simulate slow operation
  go func() {
    time.Sleep(3 * time.Second)
    ch <- "Operation complete"
  }()

  // Wait with timeout
  fmt.Println("Waiting for result...")

  select {
  case result := <-ch:
    fmt.Println("Got:", result)
  case <-time.After(2 * time.Second):
    fmt.Println("Timeout! Operation took too long")
  }

  // Multiple attempts with timeout
  ch2 := make(chan int)

  go func() {
    time.Sleep(500 * time.Millisecond)
    ch2 <- 42
  }()

  select {
  case value := <-ch2:
    fmt.Println("Received value:", value)
  case <-time.After(1 * time.Second):
    fmt.Println("Timeout waiting for value")
  }
}

