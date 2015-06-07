package main

import (
    "fmt"
    "math"
)

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    fibN2 := 0
    fibN1 := 1
    calls := 0
    result := 0
    return func() int {
        if calls == 0 {
            result = 0
        } else if calls == 1 {
            result = 1
        } else {
           result = fibN1 + fibN2
           fibN2 = fibN1
           fibN1 = result
        }
        calls++
        return result
    }
}

func fibo_main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

func old__main() {
    funct := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(funct(3, 4))
    pos, neg := adder(), adder()
    for i := 0;i<10;i++ {
        fmt.Println(pos(3),
        neg(-4))
    }
}
