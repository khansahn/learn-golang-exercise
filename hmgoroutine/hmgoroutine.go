package main

import (
	"fmt"
	"time"
	"sync"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

// channels
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 	// send sum to c
}

// range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// select
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

// sync, mutual exclusion aka mutex
/// mutex has two methods, Lock and Unlock
/// make sure only one  goroutine can access a variable at a time to avoid conflicts
//// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}
//// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mux.Unlock()
}
//// value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	go say("world")
	say("hello")

	// channels
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <- c 	// receive from c

	fmt.Println(x, y, x+y)

	// buffered channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	// range and close
	cr := make(chan int, 10)
	go fibonacci(cap(cr), cr)
	for i := range cr {
		fmt.Println(i)
	}

	// select
	cs := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<- cs)
		}
		quit <- 0
	}()
	fibonacci2(cs, quit)

	// default selection
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <- tick:
	//		fmt.Println("tick.")
	//	case <- boom:
	//		fmt.Println("BOOMMMMM")
	//		return
	//	default:
	//		fmt.Println("......")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	// sync mutex
	csm := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go csm.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println("HAHAAHHH",csm.Value("somekey"))


}
