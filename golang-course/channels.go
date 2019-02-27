package main

import "fmt"

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

func mainChannels() {
	rangeChannels()
}
