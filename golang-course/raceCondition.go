package main

import (
	"fmt"
	"runtime"
	"sync"
)

func mainRaceCondition() {
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines before: ", runtime.NumGoroutine())

	var wg1 sync.WaitGroup

	counter := 0

	const gs = 100

	wg1.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// every go func will be accessing a different v but same counter
			v := counter
			// Can run time.sleep(time.Second) or Gosched to yield the processor
			runtime.Gosched()
			v++
			counter = v
			wg1.Done()
		}()
	}

	wg1.Wait()

	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines after: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

	// There's a race condition here - the counter doesn't get updated to 100 as expected
}
