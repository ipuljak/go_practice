package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func understandingChannels() {
	// c := make(chan int) // a channel onto which I can put integers

	// c <- 42 // I put 42 into channel c

	// fmt.Println(<-c) // I'm taking out the value from c and printing it

	// This produces deadlock!
	// CHANNELS BLOCK
	// They act almost like relay races - they have to be able to hand data to something

	// Successful version 1
	c1 := make(chan int)

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)

	// Successful version 2
	c2 := make(chan int, 1)

	c2 <- 22

	fmt.Println(<-c2)

	// Unsuccessful buffer
	c3 := make(chan int, 2) // If I change 1 to 2 it'll work

	c3 <- 10
	c3 <- 11 // This is going to break because c2 is now full (space for only one) and it now blocks
	// If the channel was size 2 I could pull off the 11

	fmt.Println(<-c3)
}

func directionalChannels() {
	c := make(chan int, 2)

	// These wont work
	// c := make(chan <- int, 2)
	// c := make(chan(<-chan int, 2))

	// Receiving channel
	cr := make(<-chan int)

	// Sending channel
	cs := make(chan<- int)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Println("%T\n", c)
	fmt.Println("%T\n", cr)
	fmt.Println("%T\n", cs)
}

func usingChannels() {
	c := make(chan int)

	// send
	go fooSend(c)

	// receive
	//go barReceive(c)
	barReceive(c) // This is now blocking until the value is sent (similar to waitGroup)

	fmt.Println("About to exit...")
}

// send (SEND ONLY CHANNEL)
func fooSend(c chan<- int) {
	c <- 42
}

// receive (RECEIVE ONLY CHANNEL)
func barReceive(c <-chan int) {
	fmt.Println(<-c)
}

// Range channels
func rangeChannels() {
	c := make(chan int)

	// send
	go fooSendRange(c)

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit...")
}

// send (SEND ONLY CHANNEL)
func fooSendRange(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// Channels block
// Range over channel blocks

// Select

func selectChannels() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receieve
	receive(eve, odd, quit)

	fmt.Println("About to exit...")
}

// Onto the channels we are going to send ints
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel: ", v)
		case v := <-o:
			fmt.Println("From the odd channel: ", v)
		case v := <-q:
			fmt.Println("From the quit channel", v)
			return
		}
	}
}

// Fan In Pattern
// Have a lot of work to do. Should fan it out to as many goroutines as possible so they can all be working on something
// When we get results, we'll fan those results back into a channel and then we'll have a channel with just those results
// Uses comma ok idiom

func fanInChannels() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send1(even, odd)

	go receive1(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send1(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive1(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

// Fan out
// Taking a chunk of work, and doing everything altogether at once instead of 1 by 1
// Example: Need to process a 1000 videos
// Instead of processing them serially (one at a time), instead spawn 1000 goroutines and have each one do them

func fanOutChannels() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("About to exit...")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	// IF YOU WANT TO THROTTLE (add a max amount of goroutines to spawn)
	const goroutines = 10
	wg.Add(goroutines)
	// change next line to
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}

	// NON THROTTLING VERSION
	// for v := range c1 {
	// 	wg.Add(1)

	// 	go func(v2 int) {
	// 		c2 <- timeConsumingWork(v2)
	// 		wg.Done()
	// 	}(v)
	// }

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

func mainChannels() {
	fanOutChannels()
}
