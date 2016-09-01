// Hmm.. Let's check with forever-loop


package main

import ("log"; "sync"; "runtime")

func main() {
  var join sync.WaitGroup
  // Looks like the same effect in the case of setting runtime.GOMAXPROCS(0)
  var procs = runtime.GOMAXPROCS(runtime.NumCPU())
  log.Printf("Work in parallel: %d", procs)
  for i := 0; i < runtime.NumCPU(); i++ {
    join.Add(1)
    go func(counter_id int) {
      for {}
      defer join.Done()
    }(i)
  }
  join.Wait()
  log.Print("Done")
}
