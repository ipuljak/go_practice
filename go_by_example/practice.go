package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// GO BY EXAMPLE

const s string = "constant"

// Variables
func variables() {
	fmt.Printf("\n=== 1 - Variables ===\n\n")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

// Constants
func constants() {
	fmt.Printf("\n=== 2 - Constants ===\n\n")

	const n = 50000000000
	const d = 3e20 / n
	fmt.Println(s)
	fmt.Println(d)
}

// For
func forLoops() {
	fmt.Printf("\n=== 3 - For ===\n\n")

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}

// If / Else
func ifElse() {
	fmt.Printf("\n=== 4 - If / Else ===\n\n")

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Switch Statements
func switchStatement() {
	fmt.Printf("\n=== 5 - Switch ===\n\n")

	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("four or other")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday...")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(3)
	whatAmI("hey")
}

// Arrays
func arrays() {
	fmt.Printf("\n=== 6 - Arrays ===\n\n")

	var a [5]int // Array that holds 5 integers
	fmt.Println("Empty array is zero valued: ", a)

	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])

	fmt.Println("Length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b - ", b)

	var twoD [2][5]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D array: ", twoD)
}

// Slices (typed by their elements, not number of elements)
func slices() {
	fmt.Printf("\n=== 7 - Slices ===\n\n")

	s := make([]string, 3)
	fmt.Println("Empty slice of strings: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])

	fmt.Println("Slice length: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied slice: ", c)

	l := s[2:5]
	fmt.Println("Sliced: ", l)

	l = s[:5]
	fmt.Println("Sliced 2: ", l)

	l = s[2:]
	fmt.Println("Sliced 3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("Strings: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D Slice: ", twoD)
}

// Maps
func maps() {
	fmt.Printf("\n=== 8 - Maps ===\n\n")

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 14

	fmt.Println("Map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("Length: ", len(m))

	delete(m, "k2")
	fmt.Println("Delete: ", m)

	_, prs := m["k2"] // Indicates whether the key is present in the map
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("One line map :", n)
}

// Range - iterate over elements in a variety of data
func rangeStatement() {
	fmt.Printf("\n=== 9 - Range ===\n\n")

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // Range provides index and value
		sum += num
	}
	fmt.Println("Sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index is: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// Functions
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func functions() {
	fmt.Printf("\n=== 10 - Functions ===\n\n")

	res := plus(1, 2)
	res2 := plusPlus(1, 2, 3)

	fmt.Println("Plus: ", res)
	fmt.Println("Plus plus: ", res2)
}

// Multiple return values
func vals() (int, int) {
	return 3, 7
}

func multipleReturnValues() {
	fmt.Printf("\n=== 11 - Multiple Return Values ===\n\n")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

// Variadic functions - any number of trailing arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func variadic() {
	fmt.Printf("\n=== 12 - Variadic Functions ===\n\n")

	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5, 6, 7)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// Closures - Anonymous function - useful for when you want to define a function inline without having to name it
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	fmt.Printf("\n=== 13 - Closures ===\n\n")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

// Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func recursion() {
	fmt.Printf("\n=== 14 - Recursion ===\n\n")

	fmt.Println(fact(7))
}

// Pointers
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers() {
	fmt.Printf("\n=== 15 - Pointers ===\n\n")

	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}

// Person ...
type Person struct {
	name string
	age  int
}

// Structs
func structs() {
	fmt.Printf("\n=== 16 - Structs ===\n\n")

	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)
}

// Methods

// rect ...
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods() {
	fmt.Printf("\n=== 17 - Methods ===\n\n")

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perim())

	rp := &r
	fmt.Println("area (pointer): ", rp.area())
	fmt.Println("perimeter (pointer): ", rp.perim())
}

// Interfaces - Named collections of method signatures

type geometry interface {
	iArea() float64
	iPerim() float64
}

type iRect struct {
	width, height float64
}

type iCircle struct {
	radius float64
}

func (r iRect) iArea() float64 {
	return r.width * r.height
}

func (r iRect) iPerim() float64 {
	return 2*r.width + 2*r.height
}

func (c iCircle) iArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c iCircle) iPerim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.iArea())
	fmt.Println(g.iPerim())
}

func interfaces() {
	fmt.Printf("\n=== 18 - Interfaces ===\n\n")

	r := iRect{width: 3, height: 4}
	c := iCircle{radius: 5}

	measure(r)
	measure(c)
}

// Errors

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it lol"}
	}

	return arg + 3, nil
}

func errorHandling() {
	fmt.Printf("\n=== 19 - Errors ===\n\n")

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	fmt.Println()

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// Goroutines - lightweight thread of execution (run functions asynchronously!)

func goroutineFunc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {
	fmt.Printf("\n=== 20 - Goroutines ===\n\n")

	goroutineFunc("direct")

	// go goroutineFunc("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	//fmt.Scanln()
	fmt.Println("done")
}

// Channels - pipes that connect concurrent goroutines
// You can send values into channels from one goroutine and receive those values into another goroutine

func channels() {
	fmt.Printf("\n=== 21 - Channels ===\n\n")

	messages := make(chan string)

	go func() { messages <- "ping" }() // send a value "ping" to the messagse channel we made above

	msg := <-messages
	fmt.Println(msg)
}

// Channel buffering
// Channels are by default unbuffered
// This means they only accepts sends (chan <-) if there is a corresponding recieve (<- chan)
// Buffered channels accept a limited number of values without a corresponding receiver for those values

func buffering() {
	fmt.Printf("\n=== 22 - Buffered Channels ===\n\n")

	messages := make(chan string, 2) // Will only accept 2 strings

	messages <- "buffered"
	messages <- "channel"
	// messages <- "this shouldn't work" - this line causes a fatal error "all goroutines are asleep - deadlock"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// Channel synchronization - syncronize execution across goroutines
// Following is an example of using a blocking receive to wait for a goroutine to finish

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronization() {
	fmt.Printf("\n=== 23 - Channel Synchronization ===\n\n")

	done := make(chan bool, 1)
	go worker(done)

	<-done
}

// Channel directions
// When using channels as function parameters, you can specify if a channel
// is meant to only send or receive values.
// Increases type-safety.

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	fmt.Printf("\n=== 24 - Channel Directions ===\n\n")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// Selects - select lets you wait on multiple channel operations
func selects() {
	fmt.Printf("\n=== 25 - Select ===\n\n")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		}
	}
}

// Timeouts - needed for programs that connect to external resources

func timeouts() {
	fmt.Printf("\n=== 26 - Timeouts ===\n\n")

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}

// Non blocking channel operations - use select with default to implement non blocking sends/receives
func nonBlockingChannelOperations() {
	fmt.Printf("\n=== 27 - Non Blocking Channel Operations ===\n\n")

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message1: ", msg)
	default:
		fmt.Println("No message received...")
	}

	msg := "Hi!"

	select {
	case messages <- msg:
		fmt.Println("Sent message: ", msg)
	default:
		fmt.Println("No message sent...")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message2: ", msg)
	case sig := <-signals:
		fmt.Println("Received signal: ", sig)
	default:
		fmt.Println("No activity")
	}
}

// Closing channels
func closingChannels() {
	fmt.Printf("\n=== 28 - Closing Channels ===\n\n")

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job: ", j)
			} else {
				fmt.Println("Received all jobs!")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job: ", j)
	}

	close(jobs)

	fmt.Println("Sent all jobs!")
}

// Range over channels - iterate over values receieved from a channel
func rangeOverChannels() {
	fmt.Printf("\n=== 29 - Range Over Channels ===\n\n")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// Timers - execute code in the future or repeatedly at some interval
func timers() {
	fmt.Printf("\n=== 30 - Timers ===\n\n")

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

// Tickets - want to do something repeatedly at regular intervals
func tickers() {
	fmt.Printf("\n=== 31 - Tickers ===\n\n")

	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at: ", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// Worker Pools - runs several concurrent instances

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker: ", id, "started job - ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker: ", id, "finished job - ", j)
		results <- j * 2
	}
}

func workerPools() {
	fmt.Printf("\n=== 32 - Worker Pools ===\n\n")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}

// Rate Limiting - controlling resource utilization and maintaining quality of service
func rateLimiting() {
	fmt.Printf("\n=== 33 - Rate Limiting ===\n\n")

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Bursty request: ", req, time.Now())
	}
}

// Atomic Counters - can manage state with worker pools but also with sync/atomic package
func atomicCounters() {
	fmt.Printf("\n=== 34 - Atomic Counters ===\n\n")

	var ops uint64 // Counter for how many operations were performed

	for i := 0; i < 50; i++ { // Start 50 goroutines that each increment the counter about once a millisecond
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // We can add to the ops variable because we give it the memory address
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops) // To safely use the counter we have to load in what's in the memory address at this point
	fmt.Println("ops: ", opsFinal)
}

// Mutexes - like above can also used to access data across goroutines, but more safely
func mutexes() {
	fmt.Printf("\n=== 35 - Mutexes ===\n\n")

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	// Keeping track of read and write operations
	var readOps uint64
	var writeOps uint64

	// Start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 goroutines to simulate writes using the same pattern as above
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the above work for a second
	time.Sleep(time.Second)

	// Take and report the final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops: ", writeOpsFinal)

	// Lock the state and lets see what it looks like
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}

// Stateful Goroutines - another option of above is to use the built in synchronization
// features of goroutines and channels to achieve the same result. Aligns with Go's ideas
// of sharing memory by communicating and having each piece of data owned by only 1 goroutine

// readOp
type readOp struct {
	key  int
	resp chan int
}

// writeOp
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func statefulGoroutines() {
	fmt.Printf("\n=== 36 - Stateful Goroutines ===\n\n")

	// State will be owned by a single goroutine and all other goroutines
	// will send messages to the owning goroutine and receive replies

	var readOps uint64
	var writeOps uint64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// This goroutine owns the state
	// Repeatedly selects on the reads/writes channels, responding to requests as they arrive
	// A response is executed by first performing the requested operation and then sending a value
	// on the response channel resp to indicate success (and the desired value of reads)
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops: ", writeOpsFinal)
}

// Sorting
func sorting() {
	fmt.Printf("\n=== 37 - Sorting ===\n\n")

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints) // Check if the ints are sorted
	fmt.Println("Sorted: ", s)
}

// Sorting by functions

// Need a corresponding type to sort by a custom function
type byLength []string

// Implement the sort interface - need Len, Less, and Swap
// Len and swap are typically the same across most types
// Less is the real sorting logic here
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Essentially sorting by string size here
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func sortingByFunctions() {
	fmt.Printf("\n=== 38 - Sorting By Function ===\n\n")

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

// Panic
func panicExample() {
	fmt.Printf("\n=== 39 - Panic ===\n\n")

	// panic("a problem")

	// _, err := os.Create("/tmp/file")

	// if err != nil {
	// 	panic(err)
	// }
}

// Defer - ensure that a function call is performed later in a program's execution
// Usually for puposes of cleanup (ensure, finally used in other languages)

func createFile(p string) *os.File {
	fmt.Println("Creating file...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing...")
	f.Close()
}

func deferring() {
	fmt.Printf("\n=== 40 - Defer ===\n\n")

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

// Collection functions - perform operations on collections of data

// Index eturns the index of the given string in the given slice
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// Include returns whether the string exists in the slice
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfied the predicate f
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f
func Filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)

	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Map returns a new slice containing the results of applying the function f to each item in the original slice
func Map(vs []string, f func(string) string) []string {
	// mapped := make([]string, 0)

	// for _, v := range vs {
	// 	mapped = append(mapped, f(v))
	// }

	// return mapped

	// Another way to do this:

	mapped := make([]string, len(vs))

	for i, v := range vs {
		mapped[i] = f(v)
	}

	return mapped
}

func collectionFunctions() {
	fmt.Printf("\n=== 41 - Collection Functions ===\n\n")

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}

// String Functions
func stringFunctions() {
	fmt.Printf("\n=== 42 - String Functions ===\n\n")

	var p = fmt.Println

	p("Contains: ", strings.Contains("test", "es"))
	p("Count: ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])
}

// String Formatting

// point
type point struct {
	x, y int
}

func stringFormatting() {
	fmt.Printf("\n=== 43 - String Formatting ===\n\n")

	var p = fmt.Printf

	p1 := point{1, 2}

	p("%v\n", p1)                      // "Verbs" - format general Go values, will print an instance of our point struct
	p("%#v\n", p1)                     // "Verbs" with struct information
	p("%#v\n", p1)                     // Prints a Go syntax representation of the value (i.e. the source code snippet that would produce that value)
	p("%T\n", p1)                      // Prints the type of the variable
	p("%t\n", true)                    // Printing bool formatting
	p("%d\n", 123)                     // Standard base 10 integer formatting
	p("%b\n", 14)                      // Binary representation
	p("%c\n", 33)                      // Character corresponding to the given integer
	p("%x\n", 456)                     // Hex encoding
	p("%f\n", 78.9)                    // Floats
	p("%e\n", 123400000.0)             // Scienfitic notation
	p("%E\n", 123400000.0)             // Scientific notation 2
	p("%s\n", "\"string\"")            // Basic string printing
	p("%q\n", "\"string\"")            // Double quote strings
	p("%x\n", "hex this")              // Base 16
	p("%p\n", &p1)                     // Representation of a pointer
	p("|%6d|%6d|\n", 12, 345)          // Specifying width and precision of an integer
	p("|%6.2f|%6.2f|\n", 1.2, 3.45)    // Specifying width and precision of floats
	p("|%-6.2f|%-6.2f|\n", 1.2, 3.45)  // Left justify
	p("|%6s|%6s|\n", "foo", "b")       // Control width when formatting strings
	p("|%-6s|%-6s|\n", "foo", "b")     // Left justify use the - flag as with numbers
	s := fmt.Sprintf("a %s", "string") // Formats and returns a string without printing it anywhere
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // Format + print to io.Writers other than os.Stdout using Fprintf
}

// Regular Expressions
func regularExpressions() {
	fmt.Printf("\n=== 44 - Regular Expressions ===\n\n")

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // Direct match
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") // Compile regex first and then use it later

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

// JSON

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int `json:"page"`
	Fruits []string
}

func jsonFunctions() {
	fmt.Printf("\n=== 45 - JSON ===\n\n")

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB)

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

// Time
func timeFunctions() {
	fmt.Printf("\n=== 46 - Time Functions ===\n\n")

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

// Epoch Time
func epochTime() {
	fmt.Printf("\n=== 47 - Epoch Time ===\n\n")

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

// Time Formatting / Parsing
func timeFormattingParsing() {
	fmt.Printf("\n=== 48 - Time Formatting / Parsing ===\n\n")

	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

// Random Numbers
func randomNumbers() {
	fmt.Printf("\n=== 49 - Random Numbers ===\n\n")

	// Default number generator is deterministic - it'll produce the same numbers every time!
	// To produce proper numbers, give it a seed that changes

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

// Number Parsing
func numberParsing() {
	fmt.Printf("\n=== 50 - Number Parsing ===\n\n")

	// Parsing numbers to strings

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

// URL Parsing
func urlParsing() {
	fmt.Printf("\n=== 51 - URL Parsing ===\n\n")

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // Access the scheme

	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host) // Host contains both username and password, use this to extract both
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

// SHA1 Hashes
func sha1Hashes() {
	fmt.Printf("\n=== 52 - SHA1 Hashes ===\n\n")

	s := "sha1 this stringfsddf4"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// Base64 Encoding
func base64Encoding() {
	fmt.Printf("\n=== 53 - Base64 Encoding ===\n\n")

	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

// Reading Files

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readingFiles() {
	fmt.Printf("\n=== 54 - Reading Files ===\n\n")

	// Most basic file reading task is slurping a file's entire contents into memory
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// You'll typically need more control, so you will need the OS to first open the file
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file
	// Allow up to 5 to be read, but also note how many were actually read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// You can also seek to a known location and read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// io package provides some functions that can be helpful for reading files
	// reads like the one above can be more robustly implemented with ReadAtLeast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// No built-in rewind, but Seek(0, 0) accomplishes this
	_, err = f.Seek(0, 0)
	check(err)

	// bufio package implements a buffered reader that may be useful both for efficiency
	// with many small reads and because of the additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you are done with it
	// (Usually this would be scheduled immediately after Opening with defer)
	f.Close()
}

// Writing Files
func writingFiles() {
	fmt.Printf("\n=== 55 - Writing Files ===\n\n")

	// Dumping a string (or just bytes) into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Idiomatic to defer a close right after opening a file
	defer f.Close()

	// You can write byte slices as you would expect
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString also available
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers we saw earlier
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer
	w.Flush()
}

// Line Filters - a common type of program that reads input on stdin, processes it, and then prints some results
// grep and sed are common line filters
func lineFilters() {
	fmt.Printf("\n=== 56 - Line Filters ===\n\n")

	// An example of a line filter that writes a capitalized version of all input text

	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method
	// that advances the scanner to the next token; which is the next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text returns the current token, here the next line, from the input
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

// Command Line Arguments
func commandLineArguments() {
	fmt.Printf("\n=== 57 - Command Line Arguments ===\n\n")

	// Prints everything including how this function was run
	argsWithProg := os.Args

	// Prints just the arguments
	argsWithoutProg := os.Args[1:]

	// Prints the third argument
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// Command Line Flags
func commandLineFlags() {
	fmt.Printf("\n=== 58 - Command Line Flags ===\n\n")

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string

	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("numb: ", *numbPtr)
	fmt.Println("fork: ", *boolPtr)
	fmt.Println("svar: ", svar)
	fmt.Println("tail: ", flag.Args())
}

// Environment Variables
func environmentVariables() {
	fmt.Printf("\n=== 58 - Environment Variables ===\n\n")

	os.Setenv("FOO", "1")
	fmt.Println("FOO: ", os.Getenv("FOO"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

// Spawning Processes - sometimes our Go programs need to spawn other non-Go processes
func spawningProcesses() {
	fmt.Printf("\n=== 59 - Spawning Processes ===\n\n")

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// Executing Processes
func executingProcesses() {
	fmt.Printf("\n=== 60 - Executing Processes ===\n\n")

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}

// Signals
func signalsFunctions() {
	fmt.Printf("\n=== 61 - Signal Functions ===\n\n")

	// This program listens in on signals
	// Specifically, when we interrupt the program (CTRL-C)
	// it listens for this signals and knows we have interrupted it

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

// Exit - use os.Exit to immediately exit with a given status
func exitFunctions() {
	fmt.Printf("\n=== 62 - Exit ===\n\n")

	defer fmt.Println("!")

	// Exit with status 3
	os.Exit(3)

	// If you go run practice.go, the exit will be picked up by Go and printed
	// If you go build practice.go and then run it, you can see the status in the terminal
	// run `echo $?` after building and then running it
}

func main() {
	// variables()
	// constants()
	// forLoops()
	// ifElse()
	// switchStatement()
	// arrays()
	// slices()
	// maps()
	// rangeStatement()
	// functions()
	// multipleReturnValues()
	// variadic()
	// closures()
	// recursion()
	// pointers()
	// structs()
	// methods()
	// interfaces()
	// errorHandling()
	// goroutines()
	// channels()
	// buffering()
	// channelSynchronization()
	// channelDirections()
	// selects()
	// timeouts()
	// nonBlockingChannelOperations()
	// closingChannels()
	// rangeOverChannels()
	// timers()
	// tickers()
	// workerPools()
	// rateLimiting()
	// atomicCounters()
	// mutexes()
	// statefulGoroutines()
	// sorting()
	// sortingByFunctions()
	// panicExample()
	// deferring()
	// collectionFunctions()
	// stringFunctions()
	// stringFormatting()
	// regularExpressions()
	// jsonFunctions()
	// timeFunctions()
	// epochTime()
	// timeFormattingParsing()
	// randomNumbers()
	// numberParsing()
	// urlParsing()
	// sha1Hashes()
	// base64Encoding()
	// readingFiles()
	// writingFiles()
	// lineFilters()
	// commandLineArguments()
	// commandLineFlags()
	// environmentVariables()
	// spawningProcesses()
	// executingProcesses()
	// signalsFunctions()
	// exitFunctions()
}
