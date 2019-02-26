package main

import (
	"fmt"
	"runtime"
	"sync"
)

func mainMutex() {
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines before: ", runtime.NumGoroutine())

	var wg1 sync.WaitGroup

	counter := 0

	const gs = 100

	wg1.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			// every go func will be accessing a different v but same counter
			v := counter
			// Can run time.sleep(time.Second) or Gosched to yield the processor
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg1.Done()
		}()
	}

	wg1.Wait()

	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines after: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

	// There is a race condition because we have multiple goroutines accessing a shared variable
	// What we need is for a goroutine to HOLD onto that variable while it is using it so that nothing else can touch it until it's done
	// Mutex is all about locking access to a certain variable
}
