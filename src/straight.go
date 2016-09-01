package main

// This code runs time.Sleep synchronously.

import ("log"; "time"; "runtime")

func main () {
  log.Print("Work in straight")
  for i := 0; i < runtime.NumCPU(); i++ {
    log.Printf("Start sleeping: %d", i)
    time.Sleep(3 * time.Second)
    log.Printf("End sleeping: %d", i)
  }
  log.Print("Done")
}
