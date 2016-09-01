package main

// It runs time.sleep multiple time that is given by runtime.NumCPU()
// However, it runs asyncronously.

import ("log"; "sync"; "time"; "runtime"; "os")

func main() {
  var join sync.WaitGroup
  var procs = runtime.GOMAXPROCS(0)
  log.Printf("Work in concurrency (Not in parallel): %d", procs)
  for i := 0; i < runtime.NumCPU(); i++ {
    join.Add(1)
    go func(counter_id int) {
      log.Printf("Start sleeping: %d (pid: %d)", counter_id, os.Getpid())
      time.Sleep(3 * time.Second)
      defer log.Printf("End sleeping: %d (pid: %d)", counter_id, os.Getpid())
      defer join.Done()
    }(i)
  }
  join.Wait()
  log.Print("Done")
}
