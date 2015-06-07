package main

import (
	"fmt"
	"strconv"
	"time"
)

func say(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(strconv.Itoa(i), s)
	}
}

func _sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func fibonacciChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func concurrency_main() {
	//go say("Hello", 5)
	//say("World", 5)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 66}
	c := make(chan int)
	fmt.Println(len(a)/2, a[:len(a)/2], a[len(a)/2:])

	go _sum(a[:len(a)/2], c)
	go _sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	d := make(chan int, 3)
	d <- 1
	d <- 2
	d <- 3
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)

	fmt.Println("Fibonacci")
	f := make(chan int, 10)
	go fibonacciChan(cap(f), f)
	for i := range f {
		fmt.Println(i)
	}

	fmt.Println("FibonacciSelect")
	g := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-g)
		}
		quit <- 0
	}()
	fibonacciSelect(g, quit)

	fmt.Println("test tick")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("TICK")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
