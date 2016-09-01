package main

// It runs time.sleep multiple time that is given by runtime.NumCPU()
// However, it runs asyncronously and expected multi-process.
// As the end of result, this code doesn't run in multi-process, because go
// statuemnt runs a function asyncronously, not generating a new process.

// I felt `go func` was similar to `async def` in python :)

import ("log"; "sync"; "time"; "runtime"; "os")

func main() {
  var join sync.WaitGroup
  // Looks like the same effect in the case of setting runtime.GOMAXPROCS(0)
  var procs = runtime.GOMAXPROCS(runtime.NumCPU())
  log.Printf("Work in parallel: %d", procs)
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
