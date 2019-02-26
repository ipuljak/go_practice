package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func mainWaitGroup() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foowg()
	barwg()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()
}

func foowg() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func barwg() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
