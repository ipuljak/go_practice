package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func mainAtomic() {
	// sync.atomic provides low-level atomic memory primitives useful for implementing synchronziation algorithms
	// It's essentially one of the ways we can avoid race conditions

	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines before: ", runtime.NumGoroutine())

	var wg1 sync.WaitGroup

	var counter int64 // package atomic requires int64

	const gs = 100

	wg1.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter - ", atomic.LoadInt64(&counter))
			wg1.Done()
		}()
	}

	wg1.Wait()

	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("# Goroutines after: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

}
