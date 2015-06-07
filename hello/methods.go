package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Vertex struct {
	X, Y float64
}

type MyFloat64 float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d yo)", p.Name, p.age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func methods_main() {
	v := &Vertex{2, 3}
	fmt.Println(v.Abs())
	v.Scale(3)
	//fmt.Println(v.Abs())
	f := MyFloat64(-math.Sqrt2)
	//fmt.Println(f.Abs())

	var a Abser
	a = f
	a = v
	//a = *v

	fmt.Println(a.Abs())

	var w Writer

	w = os.Stdout
	fmt.Fprintf(w, "Hello World !\n")

	e := Person{"john", 35}
	g := Person{"Bob", 32}
	fmt.Println(e, g)

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}

	h, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
	fmt.Printf("Converted int %v\n", h)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
