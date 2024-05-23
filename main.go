package main

import (
	"fmt"
	"sync"
)

var (
  // counter is the shared resource that multiple goroutines will use to increment.
  counter int
  // mu is used to control the access of counter
  mu  sync.Mutex
  // wg is for waiting for all goroutines to finish
  wg  sync.WaitGroup
)

func increment(){
  // done is called at the end to signal that the goroutine is finished
  defer wg.Done()
  // locks the mutex means , means counter cant be accessed by other goroutines
  mu.Lock()
  counter++
  fmt.Println("in:",counter)
  // unlock it after modifying the counter
  mu.Unlock()
}

func main(){
  // here we are adding 2 goroutines on waitgroup
  wg.Add(2)
  // start 2 goroutines for incrementing the counter
  go increment()
  go increment()
  // will wait for both goroutines to finish
  wg.Wait()
  fmt.Println("counter:",counter)
}
// if we don't use mutex to lock the shared resource ie "counter" in this case .
// both goroutines will try to access the same value and change it and could cause varying final output which is not ideal.
