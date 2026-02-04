package main

import (
  "fmt"
  "time"
)

func main() {
  data := make(chan int)
  done := make(chan bool)

  // Producer
  go func() {
    for i := 1; i <= 5; i++ {
      data <- i
      time.Sleep(300 * time.Millisecond)
    }
    done <- true
  }()

  // Consumer with select in loop
  for {
    select {
    case value := <-data:
      fmt.Println("Processing:", value)
    case <-done:
      fmt.Println("All done!")
      return
    case <-time.After(1 * time.Second):
      fmt.Println("No activity, timing out")
      return
    }
  }
}

