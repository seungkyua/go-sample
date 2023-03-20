package main

import (
	"fmt"
	"runtime"
	"time"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	/* ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println("go runtine", v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println("main go runtine", v, v2) */

	/* for i := range countTo(10) {
		fmt.Println(i)
	} */

	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	done := make(chan struct{})
	nums := PlusOne(done, c)
	for num := range nums {
		fmt.Println(num)
		if num == 14 {
			break
		}
	}
	close(done)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	for num := range nums {
		fmt.Println("Consume: ", num)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}

func PlusOne(done <-chan struct{}, in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range in {
			select {
			case ch <- num + 1:
			case <-done:
				return
			}
		}
	}()
	return ch
}
